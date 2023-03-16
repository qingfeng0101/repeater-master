package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"github.com/kexirong/repeater/models"
	"github.com/kexirong/repeater/sender"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/common/model"
)

type Data struct {
	Receiver string  `json:"receiver"`
	Status   string  `json:"status"`
	Alerts   []Alert `json:"alerts"`

	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`

	ExternalURL string `json:"externalURL"`
}

// type a model.Alert
type Alert struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
}

type Message struct {
	*Data

	// The protocol version.
	Version         string `json:"version"`
	GroupKey        string `json:"groupKey"`
	TruncatedAlerts uint64 `json:"truncatedAlerts"`
}

func WebHookReceiver(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for _, alert := range message.Alerts {
		c.Logger().Info(time.Now(), alert.Labels, alert.Annotations)
		indexes := sender.Indexes{}
		for k, v := range alert.Labels {
			indexes = append(indexes, sender.Index{Label: model.LabelName(k), Value: model.LabelValue(v)})
		}
		ruleName := sender.NotifyRuleRepos().IndexCache.GetValByIndexes(indexes)
		if ruleName == "" {
			ruleName = "default"
		}
		rule := sender.NotifyRuleRepos().Get(ruleName)
		if rule == nil {
			c.Logger().Errorf("rule: %s not fount", string(ruleName))
			continue
		}
		buf := new(bytes.Buffer)
		var t, _ = template.New("alert_message").Funcs(funcMap).Parse(alertTPL)
		t.Execute(buf, alert)
		c.Logger().Info(ruleName, rule.SenderSet, rule.Contacts, rule.SenderFilter)
		severity := alert.Labels["severity"]
		for _, s := range rule.SenderSet {
			der := sender.SenderRepos().Get(s)
			if _, ok := Find(rule.SenderFilter[sender.Severity(severity)], der.Type()); ok {
				var contacts []models.Contact
				models.DB.Find(&contacts, "username in ?", rule.Contacts)
				var content json.RawMessage = buf.Bytes()
				if err := der.Send(contacts, "监控告警", "markdown", &content); err != nil {
					c.Logger().Error(err)
				}
			}
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

const alertTPL = `### 监控告警({{.Labels.environment}})  
> <font color="{{fontColor .Status .Labels.severity}}" size="5">[告警{{statusToZh .Status}}]</font>  
> **告警名称:** {{.Labels.alertname}}  
> **告警级别:** <font color="{{fontColor .Status .Labels.severity}}">{{.Labels.severity}}</font>  
{{- if .Labels.container_name}}   
> **告警容器:** {{.Labels.container_name}}{{- end}}  
{{- if .Labels.namespace}}  
> **命名空间:** {{.Labels.namespace}}{{- end}}  
> **告警详情:**    
>>  
>> 开始时间: {{timeFormat .StartsAt}}  
>> 结束时间: {{timeFormat .EndsAt}}  
>> 告警实例: {{if .Labels.service_name}}{{.Labels.service_name}}{{else}}{{.Labels.instance}}{{end}}  
>> 详细信息: {{.Annotations.message}}  `

var clorMap = map[string]string{"high": "#f0a020", "critical": "#d03050", "warning": "#af52de", "info": "#909399"}
var funcMap = template.FuncMap{
	"statusToZh": func(str string) string { return map[string]string{"firing": "触发", "resolved": "恢复"}[str] },
	"timeFormat": func(t time.Time) string {
		if t.IsZero() {
			return "-"
		}
		return t.Local().Format("2006-01-02 15:04:05")
	},
	"fontColor": func(status, severity string) string {
		if status == "resolved" {
			return "#18a058"
		}
		return clorMap[severity]
	},
}
