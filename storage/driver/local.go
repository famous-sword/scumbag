package driver

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/foundation"
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
	file, err := local.fs.Open(key)

	return file, err
}

func (local *Local) Remove(key string) error {
	err := local.fs.Remove(key)

	if err != nil {
		return err
	}

	return nil
}

func (local *Local) Sync(key, pathname string) error {
	panic("implement me")
}

func NewLocal() foundation.StorageDriver {
	local := new(Local)
	local.mount = config.String("storage.local.mount")

	fs := afero.NewOsFs()
	local.fs = afero.NewBasePathFs(fs, local.mount)

	return local
}
