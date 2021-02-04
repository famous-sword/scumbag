package stroage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"io"
	"path/filepath"
	"strings"
)

var hasher = sha256.New()

type Object struct {
	id     string
	reader io.Reader

	Name string

	Size int64
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
	hasher.Reset()
	var buffer bytes.Buffer

	object.Size, err = io.Copy(&buffer, io.TeeReader(reader, hasher))

	if err != nil {
		return err
	}

	object.reader = &buffer
	object.Hash = hex.EncodeToString(hasher.Sum(nil))
	object.Ext = strings.TrimLeft(filepath.Ext(object.Name), ".")

	return err
}

func NewObject() *Object {
	object := new(Object)
	object.id = uuid.New().String()

	return object
}
