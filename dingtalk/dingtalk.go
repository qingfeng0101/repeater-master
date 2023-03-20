package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/sender"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"errors"
)

const AccessTokenurl = "https://api.dingtalk.com/v1.0/oauth2/accessToken"
const Getuseridurl = "https://oapi.dingtalk.com/topapi/v2/user/getbymobile?access_token="
const SendMsgurl = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token="


type DingDing struct {
	name string `json:"name"`
	AppKey string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	AgentId uint `json:"agent_id"`
	TokenMsg
}

type TokenMsg struct {
	AccessToken string `json:"accessToken"`
	ExpireIn int64 `json:"expireIn"`

}

type Users struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
	Result Result `json:"result"`
}
type Results struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
}
type Result struct {
	ExclusiveAccountUseridList []string `json:"exclusive_account_userid_list"`
	Userid string `json:"userid"`
}
type ToUser struct {
	Username string `json:"username"`
	Userid string `json:"userid"`
}
func New(Name,AppKey,AppSecret string,AgentId uint) *DingDing  {
	return &DingDing{
		name: Name,
		AppKey: AppKey,
		AppSecret: AppSecret,
		AgentId: AgentId,
	}
}
// 获取钉钉基础信息方法
func (d *DingDing) GetAccessToken() error {
	b,_ := json.Marshal(*d)
	resp,err := http.Post(AccessTokenurl,"application/json;utf-8",bytes.NewReader(b))
	if err != nil{
		log.Println("GetAccessToken err: ",err)
		return err
	}
	defer resp.Body.Close()
	b,err = ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Println("ioutil.ReadAll err: ",err)
		return err
	}
	json.Unmarshal(b,d)
	d.ExpireIn += time.Now().Unix()

	return nil
}
func (d *DingDing) CheckAccessToken() error {
	if d.AccessToken == "" || d.ExpireIn - time.Now().Unix() == 0{
		var err error
		for i := 0 ;i<3;i++{
			err = d.GetAccessToken()
			if err == nil{
				return nil
			}
		}
		return err
	}
	return nil
}
func (d *DingDing) GetUid(mobile string) string{
    var Mobiles =  Mobiles{
    	Mobile: mobile,
		SupportExclusiveAccountSearch: true,
	}

	b,_ :=json.Marshal(Mobiles)

	resp,err := http.Post(Getuseridurl+d.AccessToken,"application/json;utf-8",bytes.NewReader(b))
	if err != nil{
		log.Println("Getuseridurl err: ",err)
		return ""
	}
	defer resp.Body.Close()
	b,_ = ioutil.ReadAll(resp.Body)
	var users Users
	json.Unmarshal(b,&users)
	if users.Errcode != 0{
		log.Println(users.Errmsg)
		return ""
	}

	return users.Result.Userid
}


func (d *DingDing) SendMsg(UseridList,DeptIdList string, contentType string, content *json.RawMessage) ([]byte, error) {
	//msg := Msg{
	//	Msgtype: contentType,
	//	Markdown: content,
	//}

	bm := SendMess{
		UseridList:DeptIdList,
		DeptIdList:DeptIdList,
		AgentId: d.AgentId,
	}

	//var msg interface{}
	switch contentType {
	case "markdown":
		var md MarkDown
		md.Title = "告警"
		md.Text = string(*content)
		bm.Msg.Markdown = md
		bm.Msg.Msgtype = contentType
	case "text":
		fallthrough
	default:
		bm.Msg.Msgtype= "text"
		var txt Text
		if err := json.Unmarshal(*content, &txt.Content); err != nil {
			txt.Content = string(*content)
		}
		bm.Msg.Text = txt
	}
	jmsg, err := json.Marshal(bm)

	if err != nil {
		return nil, err
	}

	//postSendmsgURL := fmt.Sprintf(SendmsgURL, d.AccessToken)

	resp, err := http.Post(SendMsgurl+d.AccessToken, "application/json;charset=utf-8", bytes.NewReader(jmsg))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}














func WecomCredentialRegist(dc models.DingCredential) {
	ding := New(dc.Name, dc.CorpID, dc.CorpSecret, dc.AgentID,)
	// name := fmt.Sprintf("%s-%s", wecom.Type(), wc.Name)
	sender.SenderRepos().Register(ding.Name(), ding)
}
func (DingDing) Type() string {
	return "dingtalk"
}

func (d DingDing) Name() string {
	return fmt.Sprintf("%s", d.name)
}
func (d DingDing) Hash() string {
	return models.WeComHash(d.AppKey, d.AgentId, d.AppSecret)
}
type NameMapUserID map[string]string
func (d *DingDing) Send(contacts []models.Contact, subject, contentType string, content *json.RawMessage) error {
	//if contentType == "html" {
	//	return errors.New("企业微信 暂时不支持html类型消息")
	//}
	var to []string
	var tod []string
	var mmui NameMapUserID
	err := d.CheckAccessToken()
	if err != nil{
		return err
	}
	for _, contact := range contacts {

		userid := d.GetUid(contact.Mobile)

		if err := json.Unmarshal([]byte(contact.DingTalk), &mmui); err != nil {
			to = append(to, contact.DingTalk)
			tod = append(tod,userid)

			continue
		}

		if uid := mmui[d.Name()]; uid != "" {
			to = append(to, uid)
			tod = append(tod,userid)


		}
	}
	//
	if len(to) == 0 {
		return nil
	}
	//
	touser := strings.Join(to,",")
	todid := strings.Join(tod,",")

	out, err := d.SendMsg(touser,todid, contentType, content)

	if err != nil {
		return err
	}
	//
	var ret Results
	json.Unmarshal(out, &ret)
	if ret.Errcode != 0 {
		return errors.New(ret.Errmsg)
	}
	return nil
}
