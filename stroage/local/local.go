package local

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/stroage"
	"github.com/spf13/afero"
	"path/filepath"
	"strings"
)

type Local struct {
	filesystem afero.Fs
}

func (local *Local) Put(object *stroage.Object) (err error) {
	hash := strings.Split(object.Id(), "-")[0]

	err = afero.WriteReader(
		local.filesystem,
		filepath.Join(hash, object.Name),
		object.Reader(),
	)

	return err
}

func (local *Local) Get(id string) *stroage.Object {
	panic("implement me")
}

func (local *Local) Delete(id string) error {
	panic("implement me")
}

func (local *Local) Remove(object *stroage.Object) error {
	panic("implement me")
}

func NewLocal() stroage.Storage {
	local := new(Local)
	fs := afero.NewOsFs()
	local.filesystem = afero.NewBasePathFs(fs, config.String("storage.local.workdir"))

	return local
}
