package driver

import "io"

type StorageDriver interface {
	Put(key string, reader io.Reader) error
	Get(key string) (io.Reader, error)
	Remove(key string) error
	Sync(key, pathname string) error
}
