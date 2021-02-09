package warp

import (
	"encoding/json"
	"github.com/famous-sword/scumbag/logger"
	"go.uber.org/zap"
)

type Meta struct {
	Version int    `json:"version"`
	Name    string `json:"name"`
	Bucket  string `json:"bucket"`
	Key     string `json:"key"`
	Size    uint64 `json:"size"`
	Hash    string `json:"hash"`
	Ext     string `json:"ext"`
}

func (meta *Meta) String() string {
	marshal, err := json.Marshal(meta)

	if err != nil {
		logger.Writter().Error("parse meta", zap.Error(err))
		return ""
	}

	return string(marshal)
}

func (meta *Meta) IncrVersion() int {
	meta.Version++

	return meta.Version
}
