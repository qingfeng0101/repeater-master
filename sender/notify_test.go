package sender

import (
	"testing"
)

func TestIndexTree(t *testing.T) {
	it := &IndexTree{root: &[]*IndexNode{}}
	for i := 0; i < len(testDataIndex); i++ {
		if err := it.Set(testDataIndex[i], testDataRuleName[i]); err != nil {
			t.Fatal(err)
		}
	}
	//验证树结构
	for i, indexes := range testDataIndex {

		var node *IndexNode
		for _, idx := range indexes {
			t.Logf("key:%s, val:%s", idx.Label, idx.Value)
			node = it.Get(idx, node)
			if node == nil {
				t.Fatal("node is nil")
			}
		}
		t.Log("================", node.Val == testDataRuleName[i])
		if node.IsLeaf() {
			t.Log("node.Val: ", node.Val)
			t.Log("GetValByIndexes: ", it.GetValByIndexes(indexes))
		}
		t.Log("================")
	}
	//测试删除
	for i1, indexes := range testDataIndex {
		it.Delete(indexes)
		t.Log("========================================================")
		for i2, indexes := range testDataIndex[i1+1:] {
			var node *IndexNode
			for _, idx := range indexes {
				t.Logf("key:%s, val:%s", idx.Label, idx.Value)
				node = it.Get(idx, node)
				if node == nil {
					t.Fatal("node is nil")
				}
			}
			t.Log("================", node.Val == testDataRuleName[i1+1+i2])
			if node.IsLeaf() {
				t.Log("node.Val: ", node.Val)
			} else {
				t.Error("node.Val not is leaf")
			}

		}
	}
	t.Error("================")

}

var testDataIndex = [][]Index{
	{
		Index{"server", "s1"},
		Index{"name", "n1"},
		Index{"page", "13"},
	},
	{
		Index{"server", "s1"},
		Index{"name", "n2"},
		Index{"tell", "t2"},
	},
	{
		Index{"server", "s1"},
		Index{"name", "n2"},
		Index{"say", "say3"},
	},
	{
		Index{"sex", "sex4"},
		Index{"height", "h4"},
		Index{"say", "say3"},
	},
}
var testDataRuleName = []RuleName{
	"rule1",
	"rule2",
	"rule3",
	"rule4",
}
