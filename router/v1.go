package router

import (
	"github.com/kexirong/repeater/handler"
	"github.com/kexirong/repeater/middleware"
	"github.com/labstack/echo/v4"
)

func AddApi(e *echo.Echo) {
	e.POST("/notify/:name", handler.AddSend, middleware.APIKey()) // 使用 指定 sets 发消息

	g := e.Group("/api", middleware.JWT())
	{
		g.GET("/credential/stmp", handler.ReadStmpCredentials)
		g.POST("/credential/stmp", handler.AddStmpCredential)
		g.DELETE("/credential/stmp/:id", handler.DeleteStmpCredential)

		g.GET("/credential/wecom", handler.ReadWecomCredentials)
		g.POST("/credential/wecom", handler.AddWecomCredential)
		g.DELETE("/credential/wecom/:id", handler.DeleteWecomCredential)

		g.GET("/credential/dingtalk", handler.ReadDingCredentials)
		g.POST("/credential/dingtalk", handler.AddDingCredential)
		g.DELETE("/credential/dingtalk/:id", handler.DeleteDingCredential)

		g.GET("/credential/lark", handler.ReadLarkCredentials)
		g.POST("/credential/lark", handler.AddLarkCredential)
		g.DELETE("/credential/lark/:id", handler.DeleteLarkCredential)

	}
	{
		g.GET("/sender", handler.ReadSenders)

		g.GET("/sender/set", handler.ReadSenderSets)
		g.GET("/sender/set/:id", handler.ReadSenderSets)
		g.POST("/sender/set", handler.AddSenderSet)
		g.PUT("/sender/set/:id", handler.UpdateSenderSet)
		g.DELETE("/sender/set/:id",handler.DeleteSenderSet)
		g.POST("/sender/send/:name", handler.AddSend) // 使用 指定 sets 发消息
		g.POST("/sender/send", handler.AddSend)       //使用default sets发消息
	}

	{

		g.GET("/notify/rule", handler.ReadNotifyRules)
		g.GET("/notify/rule/:id", handler.ReadNotifyRules)
		g.GET("/notify/rule/cache", handler.ReadNotifyRulesCache)
		g.GET("/notify/rule/index/cache", handler.ReadNotifyRulesIndexCacheTree)

		g.POST("/notify/rule", handler.AddNotifyRule)

		g.PUT("/notify/rule/:id", handler.UpdateNotifyRule)
		g.DELETE("/notify/rule/:id", handler.DeleteNotifyRule)
		g.DELETE("/notify/rule/:id/unscoped", handler.UnscopedDeleteNotifyRule)
	}
	{
		g.GET("/contact", handler.ReadContacts)
		g.GET("/contact/:id", handler.ReadContacts)
		g.POST("/contact", handler.AddContact)
		g.PUT("/contact/:id", handler.UpdateContact)
		g.POST("/contact/sync", handler.PostSyncContact)
		g.DELETE("/contact/:id", handler.DeleteContact)
	}
	{
		g.POST("/user/login", handler.Login)
		g.PATCH("/user/:username/password", handler.ChangeUserPassword)
	}
	{
		g.POST("/prometheus/alert/receiver", handler.WebHookReceiver)
		g.GET("/prometheus/datasource", handler.ReadPrometheusDataSource)
		g.GET("/prometheus/datasource/:id", handler.ReadPrometheusDataSource)
		g.POST("/prometheus/datasource", handler.AddPrometheusDataSource)
		g.PUT("/prometheus/datasource/:id", handler.UpdatePrometheusDataSource)
		g.DELETE("/prometheus/datasource/:id", handler.DeletePrometheusDataSource)
		g.GET("/prometheus/alert", handler.ReadPrometheusAlerts)
		g.GET("/prometheus/labelset/target", handler.ReadPrometheusLabelSetTarget)
		g.GET("/prometheus/severity/target", handler.ReadPrometheusRuleSeverity)
		g.GET("/prometheus/datasource/fake", handler.ReadPrometheusFakeDataSource)
		g.GET("/prometheus/datasource/fake/:id", handler.ReadPrometheusFakeDataSource)
		g.POST("/prometheus/datasource/fake", handler.AddPrometheusFakeDataSource)
		g.PUT("/prometheus/datasource/fake/:id", handler.UpdatePrometheusFakeDataSource)
		g.DELETE("/prometheus/datasource/fake/:id", handler.DeletePrometheusFakeDataSource)
	}
	{
		g.GET("/apikey", handler.ReadAPIKeys)

		g.POST("/apikey", handler.AddAPIKey)

		g.PATCH("/apikey/:id/refresh", handler.RefreshAPIKey)
		g.DELETE("/apikey/:id", handler.DeleteAPIKey)
	}
	e.Static("/assets", "dist/assets")
	e.File("/*", "dist/index.html")
}
