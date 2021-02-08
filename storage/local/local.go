package local

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/entity"
	"github.com/famous-sword/scumbag/logger"
	"github.com/famous-sword/scumbag/storage"
	"github.com/famous-sword/scumbag/storage/meta"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type Local struct {
	mount string
	fs    afero.Fs
}

func (local *Local) Put(bucket string, object *storage.Object) (err error) {
	if err = object.Validate(); err != nil {
		return err
	}

	hash := strings.Split(object.Id(), "-")[0]
	key := filepath.Join(bucket, hash, object.Name)

	err = afero.WriteReader(local.fs, key, object.Reader())

	if err != nil {
		return err
	}

	localStorage := &entity.LocalStorage{
		Uuid: object.Id(),
	}

	m := &meta.Meta{
		Version: 1,
		Bucket:  bucket,
		Name:    object.Name,
		Key:     key,
		Size:    object.Size,
		Hash:    object.Hash,
		Ext:     object.Ext,
	}

	_, err = localStorage.Create(m)

	return err
}

func (local *Local) Get(id string) (*storage.Object, error) {
	record := &entity.LocalStorage{Uuid: id}
	err := record.Load()

	if err != nil {
		logger.Writter().Error("local storage get", zap.Error(err))

		return nil, err
	}

	metas := record.MetaData()

	file, err := local.fs.Open(metas.Key)

	if err != nil {
		return nil, err
	}

	object := storage.ObjectOf(id)
	object.Name = metas.Name
	object.Hash = metas.Hash
	object.Size = metas.Size
	object.Ext = metas.Ext
	object.SetReader(file)

	return object, nil
}

func (local *Local) Delete(id string) error {
	panic("implement me")
}

func (local *Local) Remove(object *storage.Object) error {
	panic("implement me")
}

func NewLocal() storage.Storage {
	local := new(Local)
	local.mount = config.String("storage.local.mount")

	fs := afero.NewOsFs()
	local.fs = afero.NewBasePathFs(fs, local.mount)

	return local
}
