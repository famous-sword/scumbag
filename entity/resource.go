package entity

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

const (
	TypeVideo = "video"
	TypeAudio = "audio"
	TypeDoc   = "doc"
	TypeOther = "other"

	StatusCreated = "created"
)

type Meta struct {
	Output   string            `json:"output"`
	Size     uint64            `json:"size"`
	Duration uint64            `json:"duration"`
	Levels   map[string]string `json:"levels"`
}

type Resource struct {
	ID uint `gorm:"primarykey"`

	MediaId string `gorm:"size:128,unique_index"`
	Status  string `gorm:"size:16"`
	Name    string
	Type    string `gorm:"size:16"`
	Hash    string `gorm:"size:256"`
	Ext     string `gorm:"size:16"`
	Meta    string `gorm:"type:text"`
	meta    Meta   `gorm:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (r *Resource) Create(meta *Meta) (uint, error) {
	bytes, err := json.Marshal(meta)

	if err != nil {
		return 0, err
	}

	r.Meta = string(bytes)

	err = db.Create(r).Error

	return r.ID, err
}

func (r *Resource) Load() error {
	return db.Where("media_id = ?", r.MediaId).First(r).Error
}

func (r *Resource) Metas() Meta {
	return r.meta
}
