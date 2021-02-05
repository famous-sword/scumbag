package model

import (
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	ID         uint `gorm:"primarykey"`
	MediaId    string
	Ext        string `gorm:"size:16"`
	Hash       string `gorm:"size:256"`
	Size       int64
	Type       int8
	OriginName string
	Status     string `gorm:"size:16"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func CreateResource(resource *Resource) error {
	return db.Create(resource).Error
}
