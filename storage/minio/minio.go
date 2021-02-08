package minio

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/storage"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint        = config.String("storage.minio.endpoint")
	accessKeyID     = config.String("storage.minio.access_key_id")
	secretAccessKey = config.String("storage.minio.secret_access_key")
	useSSL          = config.Bool("storage.minio.ssl")
)

type Minio struct {
	client *minio.Client
}

func (m Minio) Put(bucket string, object *storage.Object) error {
	panic("implement me")
}

func (m Minio) Get(id string) (*storage.Object, error) {
	panic("implement me")
}

func (m Minio) Delete(id string) error {
	panic("implement me")
}

func (m Minio) Remove(object *storage.Object) error {
	panic("implement me")
}

func NewMinio() (storage.Storage, error) {
	var err error

	store := new(Minio)

	store.client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	return store, err
}
