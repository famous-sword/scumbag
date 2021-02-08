package ceph

import "github.com/famous-sword/scumbag/storage"

type Ceph struct{}

func (c Ceph) Put(bucket string, object *storage.Object) error {
	panic("implement me")
}

func (c Ceph) Get(id string) (*storage.Object, error) {
	panic("implement me")
}

func (c Ceph) Delete(id string) error {
	panic("implement me")
}

func (c Ceph) Remove(object *storage.Object) error {
	panic("implement me")
}

func NewCeph() storage.Storage {
	return &Ceph{}
}
