package entity

import (
	"encoding/json"
	"github.com/famous-sword/scumbag/storage/warp"
	"gorm.io/gorm"
	"time"
)

type LocalStorage struct {
	ID        uint      `gorm:"primaryKey"`
	Uuid      string    `gorm:"size:36,type:char,uniqueIndex"`
	Meta      string    `gorm:"type:text"`
	meta      warp.Meta `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ls *LocalStorage) Create(m *warp.Meta) (uint, error) {
	ls.Meta = m.String()

	err := db.Create(ls).Error

	if err != nil {
		return 0, err
	}

	return ls.ID, nil
}

func (ls *LocalStorage) Load() error {
	err := db.Where("uuid = ?", ls.Uuid).First(ls).Error

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(ls.Meta), &ls.meta)
}

func (ls *LocalStorage) MetaData() warp.Meta {
	return ls.meta
}
