package store

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/plugger"
	"github.com/famous-sword/scumbag/storage"
	"github.com/famous-sword/scumbag/storage/ceph"
	"github.com/famous-sword/scumbag/storage/local"
	"github.com/famous-sword/scumbag/storage/minio"
)

type StoragePlugger struct{}

func (p *StoragePlugger) Plug() (err error) {
	var s storage.Storage
	switch config.String("storage.driver") {
	case "minio":
		s, err = minio.NewMinio()
	case "ceph":
		s, err = ceph.NewCeph()
	default:
		s = local.NewLocal()
	}

	if err != nil {
		return err
	}

	storage.SetDriver(s)

	return err
}

func NewStoragePlugger() plugger.Plugger {
	return &StoragePlugger{}
}
