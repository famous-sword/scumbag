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

	StatusCreated          = "created"
	StatusTranscoding      = "transcoding"
	StatusTranscodeFail    = "fail"
	StatusTranscodeSuccess = "success"
)

type Meta struct {
	Output   string            `json:"output"`
	Size     uint64            `json:"size"`
	Duration uint64            `json:"duration"`
	Levels   map[string]string `json:"levels"`
}

type Resource struct {
	ID uint `gorm:"primaryKey"`

	Uuid   string `gorm:"size:36,type:char,uniqueIndex"`
	Status string `gorm:"size:16"`
	Name   string
	Type   string `gorm:"size:16"`
	Hash   string `gorm:"size:256"`
	Ext    string `gorm:"size:16"`
	Bucket string `gorm:"size:128"'`
	// for store, the string type is used
	// to avoid data being stored as blob
	Meta string `gorm:"type:text"`
	// ignore for store, readonly
	meta Meta `gorm:"-"`

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

	if len(r.Type) == 0 {
		r.Type = CastExtToType(r.Ext)
	}

	err = db.Create(r).Error

	return r.ID, err
}

func (r *Resource) Load() error {
	err := db.Where("uuid = ?", r.Uuid).First(r).Error

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(r.Meta), &r.meta)
}

func (r *Resource) Metas() Meta {
	return r.meta
}

var extensions = map[string]string{
	// videos
	"mp4":  TypeVideo,
	"avi":  TypeVideo,
	"m3u":  TypeVideo,
	"flv":  TypeVideo,
	"rm":   TypeVideo,
	"mkv":  TypeVideo,
	"mov":  TypeVideo,
	"webm": TypeVideo,
	"wmv":  TypeVideo,

	// audios
	"mp3":  TypeAudio,
	"flac": TypeAudio,
	"wav":  TypeAudio,
	"m4a":  TypeAudio,

	// docs
	"ppt":  TypeDoc,
	"pptx": TypeDoc,
	"doc":  TypeDoc,
	"docx": TypeDoc,
	"xls":  TypeDoc,
	"xlsx": TypeDoc,
	"pdf":  TypeDoc,
}

func CastExtToType(ext string) string {
	if _, has := extensions[ext]; has {
		return extensions[ext]
	}

	return TypeOther
}
