package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/kexirong/repeater/email"
	"github.com/kexirong/repeater/sender"
	"github.com/kexirong/repeater/wecom"
	"github.com/labstack/echo/v4"
)

// 此文件中的Handler 用于兼容老版本
// /sender/mail
func HandlerMailSender(c echo.Context) error {
	var pload OldPayload
	if err := c.Bind(&pload); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	r := c.Request()
	c.Logger().Infof("#sendMail# client:%s, %s-->%s, subject:%s, content:%s\n", r.RemoteAddr, pload.From, pload.To, pload.Subject, pload.Content)

	s := sender.SenderRepos().Get("email-default")
	stmp, ok := (s).(*email.SMTP)
	if !ok {
		return errors.New("wecom-default is nil, not *email.SMTP")
	}
	err := stmp.SendMail(pload.From, strings.Split(pload.To, ","), pload.Subject, pload.ContentType, []byte(pload.Content))

	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSONBlob(http.StatusOK, []byte(`{"errcode": 0, "errmsg": "ok"}`))

	}
}

// /sender/wechat
func HandlerWechatSender(c echo.Context) error {
	var payload OldPayload
	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	r := c.Request()
	c.Logger().Infof("#sendWechat# client:%s, to:%s, requestType:%s, content:%s\n", r.RemoteAddr, payload.To, payload.ContentType, payload.Content)
	s := sender.SenderRepos().Get("wecom-default")
	wc, ok := (s).(*wecom.WeCom)
	if !ok {
		return errors.New("wecom-default is nil, not *wecom.WeCom")
	}

	content := ([]byte)(fmt.Sprintf(`"%s"`, payload.Content))
	resp, err := wc.SendMsg(payload.To, "", payload.ContentType, (*json.RawMessage)(&content))

	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	ret := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}{}
	if err = json.Unmarshal(resp, &ret); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if ret.ErrCode != 0 {
		c.Logger().Error(ret.ErrMsg)
		return c.JSON(http.StatusBadRequest, ret)
	}
	c.Logger().Info(ret.ErrMsg)
	return c.JSON(http.StatusOK, ret)

}
