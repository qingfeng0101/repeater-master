package models

type LarkCredential struct {
	BaseModel
	Name       string `gorm:"column:name; type:varchar(127); default:default; not null;uniqueIndex" json:"name"`
	CorpID     string `gorm:"column:corp_id; type:varchar(127); not null" json:"corp_id" validate:"required"`
	AgentID    uint   `gorm:"column:agent_id; not null" json:"agent_id" validate:"required"`
	CorpSecret string `gorm:"column:corp_secret; type:varchar(127); not null" json:"corp_secret" validate:"required"`
}

func (LarkCredential) TableName() string {
	return "lark_credential"
}

func (dc LarkCredential) Hash() string {
	return WeComHash(dc.CorpID, dc.AgentID, dc.CorpSecret)
}
