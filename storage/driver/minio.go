package driver

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
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

func (m Minio) Put(key string, reader io.Reader) error {
	panic("implement me")
}

func (m Minio) Get(key string) (io.Reader, error) {
	panic("implement me")
}

func (m Minio) Remove(key string) error {
	panic("implement me")
}

func (m Minio) Sync(key, pathname string) error {
	panic("implement me")
}

func NewMinio() (StorageDriver, error) {
	var err error

	store := new(Minio)

	store.client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	return store, err
}
