package driver

import (
	"github.com/famous-sword/scumbag/storage/warp"
)

type StorageDriver interface {
	Put(bucket string, object *warp.Object) error
	Get(id string) (*warp.Object, error)
	Delete(id string) error
	Remove(object *warp.Object) error
}
