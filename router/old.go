package router

import (
	"net/http"

	"github.com/kexirong/repeater/handler"
	"github.com/labstack/echo/v4"
)

func AddOld(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.POST("/sender/mail", handler.HandlerMailSender)
	e.POST("/sender/wechat", handler.HandlerWechatSender)
}
