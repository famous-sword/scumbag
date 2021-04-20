package driver

import (
	"github.com/famous-sword/scumbag/foundation"
	"io"
)

type Ceph struct{}

func (c Ceph) Put(key string, reader io.Reader) error {
	panic("implement me")
}

func (c Ceph) Get(key string) (io.Reader, error) {
	panic("implement me")
}

func (c Ceph) Remove(key string) error {
	panic("implement me")
}

func (c Ceph) Sync(key, pathname string) error {
	panic("implement me")
}

func NewCeph() (foundation.StorageDriver, error) {
	return &Ceph{}, nil
}
