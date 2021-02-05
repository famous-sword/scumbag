package entity

import (
	"gorm.io/gorm"
	"time"
)

const (
	TYPE_VIDEO = "video"
	TYPE_AUDIO = "audio"
	TYPE_DOC   = "doc"
)

type Resource struct {
	ID uint `gorm:"primarykey"`

	MediaId string `gorm:"size:128,unique_index"`
	Status  string `gorm:"size:16"`
	Name    string
	Type    string `gorm:"size:16"`
	Hash    string `gorm:"size:256"`
	Size    uint64
	Ext     string `gorm:"size:16"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (r *Resource) Create() (uint, error) {
	err := db.Create(r).Error

	return r.ID, err
}
