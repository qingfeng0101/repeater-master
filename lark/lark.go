package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/sender"
	"io/ioutil"
	"net/http"
	"errors"
)



const AppAccessKey = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal"
const GetUserId = "https://open.feishu.cn/open-apis/contact/v3/users/batch_get_id?user_id_type=user_id"
const SendMsgUrl = "https://open.feishu.cn/open-apis/message/v4/batch_send/"


type LarkApp struct {
	name string `json:"name"`
	AppId string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Bodys

}
type Bodys struct {
	Status
	AppAccessToken string `json:"app_access_token"`
	Expire int `json:"expire"`
}
type UserID struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data Data `json:"data"`
}
type Data struct {
	User_List []User `json:"user_list"`
}
type User struct {
	Mobile string `json:"mobile"`
	UserId string `json:"user_id"`
	Mobiles []string `json:"mobiles"`

}


type Status struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func NewLark(Name,AppId,AppSecret string) *LarkApp  {
	return &LarkApp{
		name: Name,
		AppId: AppId,
		AppSecret: AppSecret,
	}
}
func (l *LarkApp) GetAccessKey() error {
	b ,_ := json.Marshal(l)
	resp , err := http.Post(AppAccessKey,"application/json; charset=utf-8",bytes.NewReader(b))
	if err != nil{
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	b,_ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b,l)
	if err != nil{
		return err
	}
	return nil
}
func (l *LarkApp) GetUserId(Mobiles []string)  []string {
	var Users =  &User{
		Mobiles: Mobiles,
	}

	b ,_ := json.Marshal(Users)
	re,err := http.NewRequest("POST",GetUserId,bytes.NewReader(b))
	if err != nil{
		fmt.Println("http.NewRequest err: ",err)
		return nil
	}

	re.Header.Add("Content-Type", "application/json")
	re.Header.Add("Authorization","Bearer "+ l.AppAccessToken )
	client := &http.Client{}
	resp, err := client.Do(re)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	var UserID UserID
	b,_ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(b,&UserID)
	var userid []string
	for _,user := range UserID.Data.User_List{
		userid = append(userid,user.UserId)
	}

	return userid
}

func (l *LarkApp) SendMsg(userid []string,contentType string, content *json.RawMessage) ([]byte, error){
	var Element Element
	var SendMsg SendMsg
	switch contentType {
	case "markdown":
		var Markdown =  Markdown{
			Tag: contentType,
			Content: string(*content),
		}
		Element.Elements = append(Element.Elements,Markdown)
		SendMsg.MsgType = "interactive"
		SendMsg.Card = Element
	case "text":
		fallthrough
	default:
		SendMsg.MsgType = contentType
		SendMsg.Content.Text = string(*content)
	}
	SendMsg.UserIds = userid
	b,_ := json.Marshal(SendMsg)
	re,err := http.NewRequest("POST",SendMsgUrl,bytes.NewReader(b))
	if err != nil{
		fmt.Println("http.NewRequest SendMsgUrl err: ",err)
		return nil,err
	}

	re.Header.Add("Content-Type", "application/json")
	re.Header.Add("Authorization","Bearer "+ l.AppAccessToken )
	client := &http.Client{}
	resp, err := client.Do(re)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)


}


func LarkCredentialRegist(dc models.LarkCredential) {
	lark:= NewLark(dc.Name, dc.CorpID, dc.CorpSecret)
	// name := fmt.Sprintf("%s-%s", wecom.Type(), wc.Name)
	sender.SenderRepos().Register(lark.Name(), lark)
}
func (LarkApp) Type() string {
	return "lark"
}

func (d LarkApp) Name() string {
	return fmt.Sprintf("%s", d.name)
}
func (d LarkApp) Hash() string {
	return models.LarkHash(d.AppId, d.AppSecret)
}
type NameMapUserID map[string]string
func (d *LarkApp) Send(contacts []models.Contact, subject, contentType string, content *json.RawMessage) error {
	//if contentType == "html" {
	//	return errors.New("企业微信 暂时不支持html类型消息")
	//}

	var Mobiles []string
	err := d.GetAccessKey()
	if err != nil{
		fmt.Println("send GetAccessKey err: ",err)
		return err
	}
	for _, contact := range contacts {
		Mobiles = append(Mobiles,contact.Mobile)

	}
	userid := d.GetUserId(Mobiles)
	//
	if len(userid) == 0 {
		return nil
	}
	//

	out, err := d.SendMsg(userid, contentType, content)

	if err != nil {
		return err
	}
	//
	var ret Status
	json.Unmarshal(out, &ret)
	if ret.Code != 0 {
		return errors.New(ret.Msg)
	}
	return nil
}








