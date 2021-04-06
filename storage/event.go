package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/famous-sword/scumbag/entity"
	"github.com/google/uuid"
	"io"
	"path/filepath"
	"strings"
)

var hasher = sha256.New()

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
