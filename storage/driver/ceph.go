package driver

import (
	"github.com/famous-sword/scumbag/storage/warp"
)

type Ceph struct{}

func (c Ceph) Put(bucket string, object *warp.Object) error {
	panic("implement me")
}

func (c Ceph) Get(id string) (*warp.Object, error) {
	panic("implement me")
}

func (c Ceph) Delete(id string) error {
	panic("implement me")
}

func (c Ceph) Remove(object *warp.Object) error {
	panic("implement me")
}

func NewCeph() (StorageDriver, error) {
	return &Ceph{}, nil
}
