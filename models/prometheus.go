package models

import "gorm.io/datatypes"

type NotifyRule struct {
	BaseModel
	Name         string         `gorm:"column:name; type:varchar(127); default:default; not null;uniqueIndex" json:"name"`
	LabelSetList datatypes.JSON `gorm:"column:label_set_list; not null" json:"label_set_list" validate:"required"` //锚定标签
	SenderFilter datatypes.JSON `gorm:"column:sender_filter; not null" json:"sender_filter" validate:"required"`   //通知发送过滤
	Comment      string         `gorm:"column:comment; type:longtext" json:"comment"`                              //备注
	Contacts     []Contact      `gorm:"many2many:notify_rule_x_contact" json:"contacts"`                           //通知联系人
	SenderSetID  *uint          `gorm:"column:sender_set_id" json:"sender_set_id"`
	SenderSet    *SenderSet     `json:"sender_set"` //通知发送器
}

func (NotifyRule) TableName() string {
	return "notify_rule"
}

type PrometheusDataSource struct {
	BaseModel
	Name    string `gorm:"column:name; type:varchar(127); not null;uniqueIndex" json:"name" validate:"required"`
	BaseUrl string `gorm:"column:base_url; type:varchar(127); not null" json:"base_url" validate:"required"`
	Alerts  string `gorm:"column:alerts; type:varchar(127)" json:"alerts" validate:"required"`   // /api/v1/alerts
	Targets string `gorm:"column:targets; type:varchar(127)" json:"targets" validate:"required"` // /api/v1/targets
	Rules   string `gorm:"column:rules; type:varchar(127)" json:"rules" validate:"required"`     // /api/v1/Rules
	Comment string `gorm:"column:comment; type:longtext" json:"comment"`                         //备注
}

func (PrometheusDataSource) TableName() string {
	return "prometheus_data_source"
}

type PrometheusFakeDataSource struct {
	BaseModel
	Name         string         `gorm:"column:name; type:varchar(127); not null;uniqueIndex" json:"name" validate:"required"`
	LabelSetList datatypes.JSON `gorm:"column:label_set_list; not null" json:"label_set_list" validate:"required"` //标签集合
	SeveritySet  datatypes.JSON `gorm:"column:severity_set; not null" json:"severity_set" validate:"required"`     //严重程度列表
	Comment      string         `gorm:"column:comment; type:longtext" json:"comment"`                              //备注
}

func (PrometheusFakeDataSource) TableName() string {
	return "prometheus_fake_data_source"
}
