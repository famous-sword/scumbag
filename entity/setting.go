package entity

import (
	"time"
)

type Setting struct {
	ID        uint `gorm:"primarykey"`
	Key       string
	Data      string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
