package models

import "gorm.io/datatypes"

type SenderSet struct {
	BaseModel
	Name string         `gorm:"column:name; type:varchar(127);default:default; not null;uniqueIndex" json:"name"`
	Sets datatypes.JSON `gorm:"column:sets; not null" json:"sets" validate:"required"`
}

func (SenderSet) TableName() string {
	return "sender_set"
}
