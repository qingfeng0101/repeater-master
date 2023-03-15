package wecom

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/sender"
)

const (
	AccTokenURL            = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	SendmsgURL             = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
	getDepartmentURL       = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s"
	getUserURL             = "https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%d&fetch_child=%d"
	TokeExpSec       int64 = 7200
)

/*
   "errcode": 0，
   "errmsg": "ok"，
   "access_token": "accesstoken000001",
   "expires_in": 7200
*/
type extend struct {
	Result
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type WeCom struct {
	name,
	corpID string
	agentID    uint
	corpSecret string
	extend
}

var TLSClient *http.Client

func init() {
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(wxRootPEM))
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}

	TLSClient = &http.Client{Transport: tr}

}

func New(name, corpID string, agentID uint, corpSecret string) *WeCom {
	return &WeCom{
		name:       name,
		corpID:     corpID,
		agentID:    agentID,
		corpSecret: corpSecret,
	}
}

func (w *WeCom) GetAccessToken() error {
	getAccTokenURL := fmt.Sprintf(AccTokenURL, w.corpID, w.corpSecret)

	resp, err := TLSClient.Get(getAccTokenURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&w.extend)
	if err != nil {
		return err
	}

	if w.ErrCode != 0 {
		return fmt.Errorf("get WeChat Access Token error: %s", w.ErrMsg)
	}

	w.ExpiresIn += time.Now().Unix()
	return nil
}

func (w *WeCom) checkAccessToken() error {
	if w.AccessToken == "" || w.ExpiresIn-time.Now().Unix() <= 0 {
		var err error
		for i := 0; i < 3; i++ {
			err = w.GetAccessToken()
			if err == nil {
				return nil
			}
		}
		return err
	}
	return nil
}

func (w *WeCom) SendMsg(touser, toparty, contentType string, content *json.RawMessage) ([]byte, error) {
	bm := BaseMsg{
		ToUser:  touser,
		ToParty: toparty,
		MsgType: contentType,
		AgentID: w.agentID,
	}
	var msg interface{}
	switch contentType {
	case "textcard":
		var tc TextCard
		if err := json.Unmarshal(*content, &tc); err != nil {
			return nil, err
		}
		msg = struct {
			BaseMsg
			TextCard TextCard `json:"textcard"`
		}{
			bm,
			tc,
		}

	case "markdown":
		var md Markdown
		if err := json.Unmarshal(*content, &md.Content); err != nil {
			md.Content = string(*content)
		}
		msg = struct {
			BaseMsg
			Markdown Markdown `json:"markdown"`
		}{
			bm,
			md,
		}
	case "text":
		fallthrough
	default:
		bm.MsgType = "text"
		var txt Text
		if err := json.Unmarshal(*content, &txt.Content); err != nil {
			txt.Content = string(*content)
		}
		msg = struct {
			BaseMsg
			Text Text `json:"text"`
		}{
			bm,
			txt,
		}
	}

	if err := w.checkAccessToken(); err != nil {
		return nil, err
	}

	jmsg, err := json.Marshal(msg)

	if err != nil {
		return nil, err
	}

	postSendmsgURL := fmt.Sprintf(SendmsgURL, w.AccessToken)

	resp, err := TLSClient.Post(postSendmsgURL, "application/json;charset=utf-8", bytes.NewReader(jmsg))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Send implement interface Sender
func (WeCom) Type() string {
	return "wecom"
}

func (w WeCom) Name() string {
	return fmt.Sprintf("%s", w.name)
}

func (w WeCom) Hash() string {
	return models.WeComHash(w.corpID, w.agentID, w.corpSecret)
}

type NameMapUserID map[string]string

func (w *WeCom) Send(contacts []models.Contact, subject, contentType string, content *json.RawMessage) error {

	if contentType == "html" {
		return errors.New("企业微信 暂时不支持html类型消息")
	}
	var to []string
	var mmui NameMapUserID
	for _, contact := range contacts {
		if err := json.Unmarshal([]byte(contact.WeCom), &mmui); err != nil {
			to = append(to, contact.WeCom)
			continue
		}

		if uid := mmui[w.Name()]; uid != "" {
			to = append(to, uid)
		}
	}

	if len(to) == 0 {
		return nil
	}

	out, err := w.SendMsg(strings.Join(to, "|"), "", contentType, content)

	if err != nil {
		return err
	}

	var ret Result
	json.Unmarshal(out, &ret)
	if ret.ErrCode != 0 {
		return errors.New(ret.ErrMsg)
	}
	return nil
}

func WecomCredentialRegist(wc models.WecomCredential) {
	wecom := New(wc.Name, wc.CorpID, wc.AgentID, wc.CorpSecret)
	// name := fmt.Sprintf("%s-%s", wecom.Type(), wc.Name)
	sender.SenderRepos().Register(wecom.Name(), wecom)
}

type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parentid"`
}
type respDepartment struct {
	Result
	Departments []Department `json:"department"`
}

func (w *WeCom) GetDepartmentList() ([]Department, error) {
	if err := w.checkAccessToken(); err != nil {
		return nil, err
	}
	GetDepartmentURL := fmt.Sprintf(getDepartmentURL, w.AccessToken)

	resp, err := TLSClient.Get(GetDepartmentURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var department respDepartment
	err = decoder.Decode(&department)
	if err != nil {
		return nil, err
	}

	if department.ErrCode != 0 {
		return nil, fmt.Errorf("get WeCom department list error: %s", department.ErrMsg)
	}
	return department.Departments, nil

}

type User struct {
	UserID string `json:"userid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type respUser struct {
	Result
	Users []User `json:"userlist"`
}

func (w *WeCom) GetUserList(departmentID int, fetchChild bool) ([]User, error) {

	var _fetchChild int
	if fetchChild {
		_fetchChild = 1
	}
	if err := w.checkAccessToken(); err != nil {
		return nil, err
	}
	GetUserURL := fmt.Sprintf(getUserURL, w.AccessToken, departmentID, _fetchChild)

	resp, err := TLSClient.Get(GetUserURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var user respUser
	err = decoder.Decode(&user)
	if err != nil {
		return nil, err
	}

	if user.ErrCode != 0 {
		return nil, fmt.Errorf("get WeChat Access Token error: %s", user.ErrMsg)
	}
	return user.Users, nil
}
