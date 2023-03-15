package sender

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/kexirong/repeater/config"
	"github.com/kexirong/repeater/models"
	"github.com/prometheus/common/model"
)

var notifyRuleRepos *NotifyRuleCache

func NotifyRuleRepos() *NotifyRuleCache {
	return notifyRuleRepos
}

// type LabelSet map[string]string
type RuleName string
type NotifyRuleCache struct {
	mtx        sync.Mutex
	cache      map[RuleName]*Rule
	IndexCache *IndexTree
}

type SenderSet []string
type ContactUsername string
type Severity string
type SenderFilter map[Severity][]string

type Rule struct {
	LabelSetList []model.LabelSet //规则使用LabelSet锚定
	SenderSet    SenderSet        //通知发送器集合
	Contacts     []ContactUsername
	SenderFilter SenderFilter
}

func (n *NotifyRuleCache) Regist(rule *models.NotifyRule) error {

	drule := Rule{}

	if err := json.Unmarshal(rule.LabelSetList, &drule.LabelSetList); err != nil {
		return err
	}

	if err := json.Unmarshal(rule.SenderSet.Sets, &drule.SenderSet); err != nil {
		return err
	}
	if err := json.Unmarshal(rule.SenderFilter, &drule.SenderFilter); err != nil {
		return err
	}
	for _, contact := range rule.Contacts {
		drule.Contacts = append(drule.Contacts, ContactUsername(contact.Username))
	}

	return n.Set(RuleName(rule.Name), &drule)

}

func (n *NotifyRuleCache) Update(name RuleName, rule *Rule) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	crule, ok := n.cache[name]
	if !ok {
		return errors.New(string(name) + ": not exist")
	}

	if len(rule.LabelSetList) > 0 {
		//更新LabelSetList
		n.delete(name)
		err1 := n.makeIndexByLabelSetList(rule.LabelSetList, name)
		if err1 == nil {
			n.cache[name] = rule
			return nil
		}
		//错误后回退
		if err2 := n.makeIndexByLabelSetList(crule.LabelSetList, name); err2 != nil {
			return errors.New("unknown exception, need to rebuild the cache")
		}
		n.cache[name] = crule
		return err1
	}
	//未更新LabelSetList 时
	crule.Contacts = rule.Contacts
	crule.SenderFilter = rule.SenderFilter
	crule.SenderSet = rule.SenderSet

	return nil
}

func (n *NotifyRuleCache) makeIndexByLabelSetList(labelSetList []model.LabelSet, name RuleName) error {

	errMap := ErrMap{}
	added := []Indexes{}
	for _, labelSet := range labelSetList {
		indexes := Indexes{}
		for k, v := range labelSet {
			indexes = append(indexes, Index{k, v})
		}

		if err := n.IndexCache.Set(indexes, name); err != nil {
			errMap[labelSet.String()] = err
			continue
		}
		added = append(added, indexes)

	}
	if len(errMap) != 0 {
		for _, indexes := range added {
			n.IndexCache.Delete(indexes)
		}
		return errors.New(errMap.String())
	}

	return nil

}

func (n *NotifyRuleCache) Set(name RuleName, rule *Rule) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	if _, ok := n.cache[name]; ok {
		return errors.New(string(name) + ": exist")
	}
	n.cache[name] = rule

	if err := n.makeIndexByLabelSetList(rule.LabelSetList, name); err != nil {
		delete(n.cache, name)
		return err
	}
	return nil
}

type ErrMap map[string]error

func (e ErrMap) String() string {
	emtrs := make([]string, 0, len(e))
	for s, e := range e {
		emtrs = append(emtrs, fmt.Sprintf("%s:%q", s, e))
	}
	return strings.Join(emtrs, ", ")
}

func (n *NotifyRuleCache) Cache() map[RuleName]*Rule {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	return n.cache
}

func (n *NotifyRuleCache) Len() int {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	return len(n.cache)
}

func (n *NotifyRuleCache) Get(name RuleName) *Rule {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	return n.cache[name]
}

func (n *NotifyRuleCache) Delete(name RuleName) {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	n.delete(name)
}

func (n *NotifyRuleCache) delete(name RuleName) {
	rule := n.cache[name]
	if rule == nil {
		return
	}
	delete(n.cache, name)
	for _, labelSet := range rule.LabelSetList {
		indexes := Indexes{}
		for k, v := range labelSet {
			indexes = append(indexes, Index{k, v})
		}

		n.IndexCache.Delete(indexes)
	}
}

func NewNotifyRuleCache() *NotifyRuleCache {
	return &NotifyRuleCache{
		cache:      make(map[RuleName]*Rule),
		IndexCache: &IndexTree{root: &[]*IndexNode{}},
	}
}

var subSeq map[model.LabelName]int

func init() {
	notifyRuleRepos = NewNotifyRuleCache()
	subSeq = make(map[model.LabelName]int)
	conf := config.Get()
	if conf != nil {
		for i, v := range conf.LabelOrder {
			subSeq[v] = i
		}
	}
}

type IndexTree struct {
	mutex sync.RWMutex
	root  *[]*IndexNode
}

func (t *IndexTree) Root() *IndexNode {
	return &IndexNode{
		Key:   Index{"root", ""},
		Nodes: t.root,
	}
}

func (t *IndexTree) ToMap() map[string]interface{} {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.Root().toMap()
}

func (t *IndexTree) GetValByIndexes(indexes Indexes) RuleName {

	t.mutex.Lock()
	defer t.mutex.Unlock()
	sort.Sort(indexes)
	root := t.Root()
	var nRoot *IndexNode
	for i := 0; i < len(indexes); i++ {
		nRoot = t.get(indexes[i], root)
		if nRoot == nil {
			continue
		}
		root = nRoot
		indexes = indexes[i+1:]
		i = -1
	}
	if root != nil && root.IsLeaf() {
		return root.Val
	}
	return ""
}

func (t *IndexTree) Get(index Index, root *IndexNode) *IndexNode {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.get(index, root)
}

func (t *IndexTree) get(index Index, node *IndexNode) *IndexNode {

	ns := t.root
	if node != nil {
		ns = node.Nodes
	}
	for _, node := range *ns {
		if node.Key.Label == index.Label && node.Key.Value == index.Value {
			return node
		}
	}
	return nil
}

func (t *IndexTree) Set(indexes Indexes, val RuleName) error {
	if len(indexes) == 0 {
		return errors.New("indexes length must be greater than zore")
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()
	sort.Sort(indexes)
	root := t.Root()

	for _, index := range indexes {
		nNode := t.get(index, root)

		if nNode == nil {
			nNode = &IndexNode{
				Key:   index,
				Nodes: &[]*IndexNode{},
			}
			*root.Nodes = append(*root.Nodes, nNode)
		}
		root = nNode
	}
	if root.IsLeaf() {
		return errors.New("indexes already exists")
	}
	root.Val = val
	return nil
}

func (t *IndexTree) Delete(indexes Indexes) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if len(indexes) == 0 {
		return nil
	}
	sort.Sort(indexes)
	root := t.Root()

	return t.delete(indexes, root)
}

func (t *IndexTree) delete(indexes Indexes, root *IndexNode) error {
	index := indexes[0]
	node := t.get(index, root)
	if node == nil {
		return errors.New("not fount")
	}
	if len(*node.Nodes) > 0 {
		if err := t.delete(indexes[1:], node); err != nil { //递归
			return err
		}
	}
	if len(*node.Nodes) == 0 {
		root.delete(index) //删除叶子节点
		return nil
	}
	return nil
}

type Index struct {
	Label model.LabelName
	Value model.LabelValue
}

func (i Index) String() string {
	if i.Value == "" {
		return string(i.Label)
	}
	return fmt.Sprintf("%s=%s", i.Label, i.Value)
}

type IndexNode struct {
	Key   Index         //节点键，label
	Nodes *[]*IndexNode //子节点
	Val   RuleName      //叶子节点, 保存规则名
}

func (n *IndexNode) toMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["key"] = n.Key.String()
	if n.IsLeaf() {
		m["value"] = n.Val
	}
	if len(*n.Nodes) > 0 {
		children := []map[string]interface{}{}
		for _, node := range *n.Nodes {
			children = append(children, node.toMap())
		}
		m["children"] = children
	}
	return m
}

func (n *IndexNode) delete(key Index) {
	for i := 0; i < len(*n.Nodes); i++ {
		node := (*n.Nodes)[i]
		if node.Key.Label == key.Label && node.Key.Value == key.Value {
			*n.Nodes = append((*n.Nodes)[:i], (*n.Nodes)[i+1:]...)
			return
		}
	}
}

func (n IndexNode) IsLeaf() bool {
	return n.Val != ""
}

type Indexes []Index

func (s Indexes) Len() int { return len(s) }
func (s Indexes) Less(i, j int) bool {
	return labelCompare(s[i].Label, s[j].Label)
}
func (s Indexes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// 0 if a == b, -1 if a < b, and +1 if a > b.
func labelCompare(s1, s2 model.LabelName) bool {
	p1 := subSeq[s1]
	p2 := subSeq[s2]
	switch {
	case p1 > p2:
		return true
	case p1 < p2:
		return false
	default:
		return strings.Compare(string(s1), string(s2)) < 0
	}
}
