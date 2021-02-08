package driver

import (
	"github.com/famous-sword/scumbag/resource"
)

type StorageDriver interface {
	Put(bucket string, object *resource.Object) error
	Get(id string) (*resource.Object, error)
	Delete(id string) error
	Remove(object *resource.Object) error
}
