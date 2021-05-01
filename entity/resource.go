package entity

import (
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

type Resource struct {
	ID uint `gorm:"primaryKey"`

	Uuid   string `gorm:"size:36,type:char,uniqueIndex"`
	Name   string
	Type   string `gorm:"size:16"`
	Hash   string `gorm:"size:256"`
	Size   int64
	Ext    string `gorm:"size:16"`
	Key    string
	Bucket string `gorm:"size:128"'`
	Status string `gorm:"size:16"`
	// for store, the string type is used

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ResourceRepository struct {
	resource *Resource
}

func (r *ResourceRepository) FindByUuid(id string) *Resource {
	db.Where("uuid = ?", id).Take(r.resource)

	return r.resource
}

func (r *ResourceRepository) CastExtToType(ext string) string {
	if _, has := extensions[ext]; has {
		return extensions[ext]
	}

	return TypeOther
}

func (r *ResourceRepository) Save() error {
	return db.Create(r.resource).Error
}

func (r *ResourceRepository) Uuid(uuid string) *ResourceRepository {
	r.resource.Uuid = uuid

	return r
}

func (r *ResourceRepository) Status(status string) *ResourceRepository {
	r.resource.Status = status

	return r
}

func (r *ResourceRepository) Name(name string) *ResourceRepository {
	r.resource.Name = name

	return r
}

func (r *ResourceRepository) Type(typeName string) *ResourceRepository {
	r.resource.Type = typeName

	return r
}

func (r *ResourceRepository) Hash(hash string) *ResourceRepository {
	r.resource.Hash = hash

	return r
}

func (r *ResourceRepository) Ext(ext string) *ResourceRepository {
	r.resource.Ext = ext
	r.resource.Type = r.CastExtToType(ext)

	return r
}

func (r *ResourceRepository) Bucket(bucket string) *ResourceRepository {
	r.resource.Bucket = bucket

	return r
}

func (r *ResourceRepository) Key(key string) *ResourceRepository {
	r.resource.Key = key

	return r
}

func (r *ResourceRepository) Size(size int64) *ResourceRepository {
	r.resource.Size = size

	return r
}

func NewResourceRepository() *ResourceRepository {
	r := &ResourceRepository{}
	r.resource = &Resource{}

	return r
}
