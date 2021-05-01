package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/famous-sword/scumbag/entity"
	"github.com/famous-sword/scumbag/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"path/filepath"
	"strings"
)

var hasher = sha256.New()

// Put storage file to database, then sync to storage driver
func Put(bucket string, name string, reader io.Reader) error {
	var buffer bytes.Buffer
	hasher.Reset()
	size, err := io.Copy(&buffer, io.TeeReader(reader, hasher))

	if err != nil {
		return err
	}

	hash := hex.EncodeToString(hasher.Sum(nil))
	id := uuid.New().String()
	ext := filepath.Ext(name)
	key := filepath.Join(bucket, id) + ext

	Driver().Put(key, reader)

	repository := entity.NewResourceRepository()

	repository.
		Uuid(id).
		Name(name).
		Size(size).
		Hash(hash).
		Key(key).
		Ext(strings.TrimLeft(ext, ".")).
		Status(entity.StatusCreated).
		Bucket(bucket)

	return repository.Save()
}

// Read read file from database, then read
// binary from driver
func Read(id string) (io.Reader, error) {
	resource := entity.NewResourceRepository().FindByUuid(id)

	if len(resource.Key) == 0 {
		logger.Writer().Error("file {} not exists", zap.String("{}", resource.Key))
		return nil, errors.New(fmt.Sprintf("resource #[%s] not exists", id))
	}

	reader, err := Driver().Get(resource.Key)

	if err != nil {
		return nil, err
	}

	return reader, nil
}
