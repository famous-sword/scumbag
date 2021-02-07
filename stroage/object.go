package stroage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"hash"
	"io"
	"path/filepath"
	"strings"
	"sync"
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

func (object *Object) Read(reader io.Reader) (err error) {
	hasher := hasherPool.Get().(hash.Hash)
	hasher.Reset()
	defer hasherPool.Put(hasher)

	var buffer bytes.Buffer

	size, err := io.Copy(&buffer, io.TeeReader(reader, hasher))

	if err != nil {
		return err
	}

	object.reader = &buffer
	object.Hash = hex.EncodeToString(hasher.Sum(nil))
	object.Ext = strings.TrimLeft(filepath.Ext(object.Name), ".")
	object.Size = uint64(size)

	return err
}

func NewObject() *Object {
	object := new(Object)
	object.id = uuid.New().String()

	return object
}

func LoadObject(id string) *Object {
	object := new(Object)
	object.id = id

	return object
}
