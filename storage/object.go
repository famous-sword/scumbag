package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"hash"
	"io"
	"path/filepath"
	"strings"
	"sync"
)

var (
	errHashEmpty = errors.New("object hash is empty")
	errEmptyFile = errors.New("object size is 0, signify a empty file")
)

var hasherPool = sync.Pool{
	New: func() interface{} {
		return sha256.New()
	},
}

type Object struct {
	id     string
	reader io.Reader

	Name string
	Size uint64
	Hash string
	Ext  string
}

func (object *Object) Id() string {
	return object.id
}

func (object *Object) Reader() io.Reader {
	return object.reader
}

func (object *Object) SetReader(reader io.Reader) {
	object.reader = reader
}

func (object *Object) Read(reader io.Reader) (err error) {
	hasher, ok := hasherPool.Get().(hash.Hash)

	if !ok {
		return errors.New("get hasher fail")
	}

	hasher.Reset()
	defer hasherPool.Put(hasher)

	var buffer bytes.Buffer

	size, err := io.Copy(&buffer, io.TeeReader(reader, hasher))

	if err != nil {
		return err
	}

	object.SetReader(&buffer)
	object.Hash = hex.EncodeToString(hasher.Sum(nil))
	object.Ext = strings.TrimLeft(filepath.Ext(object.Name), ".")
	object.Size = uint64(size)

	return err
}

func (object *Object) Validate() error {
	if len(object.Hash) == 0 {
		return errHashEmpty
	}

	if object.Size == 0 {
		return errEmptyFile
	}

	return nil
}

func NewObject() *Object {
	object := new(Object)
	object.id = uuid.NewString()

	return object
}

func ObjectOf(id string) *Object {
	object := new(Object)
	object.id = id

	return object
}
