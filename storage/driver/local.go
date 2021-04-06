package driver

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/spf13/afero"
	"io"
)

type Local struct {
	mount string
	fs    afero.Fs
}

func (local *Local) Put(key string, reader io.Reader) error {
	return afero.WriteReader(local.fs, key, reader)
}

func (local *Local) Get(key string) (io.Reader, error) {
	panic("implement me")
}

func (local *Local) Remove(key string) error {
	panic("implement me")
}

func (local *Local) Sync(key, pathname string) error {
	panic("implement me")
}

func NewLocal() StorageDriver {
	local := new(Local)
	local.mount = config.String("storage.local.mount")

	fs := afero.NewOsFs()
	local.fs = afero.NewBasePathFs(fs, local.mount)

	return local
}
