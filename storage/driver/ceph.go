package driver

import (
	"github.com/famous-sword/scumbag/resource"
)

type Ceph struct{}

func (c Ceph) Put(bucket string, object *resource.Object) error {
	panic("implement me")
}

func (c Ceph) Get(id string) (*resource.Object, error) {
	panic("implement me")
}

func (c Ceph) Delete(id string) error {
	panic("implement me")
}

func (c Ceph) Remove(object *resource.Object) error {
	panic("implement me")
}

func NewCeph() (StorageDriver, error) {
	return &Ceph{}, nil
}
