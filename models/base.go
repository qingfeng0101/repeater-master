package models

import (
	"gorm.io/gorm"

	"gorm.io/plugin/soft_delete"
)

var DB *gorm.DB

type BaseModel struct {
	ID        uint                  `gorm:"column:id; primary_key" json:"id"`
	CreatedAt uint                  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt uint                  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
