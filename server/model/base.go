package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	ID uuid.UUID `gorm:"type:varchar(255);primarykey" json:"id"`
	TimeStamp
}

type TimeStamp struct {
	CreateTime int64 `gorm:"autoCreateTime;index" json:"created_at"`
	UpdateTime int64 `gorm:"autoUpdateTime;index" json:"updated_at"`
}

type SoftDelete struct {
	DeleteTime soft_delete.DeletedAt `gorm:"index" json:"deleted_at"`
}
