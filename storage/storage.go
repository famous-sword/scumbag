package storage

type Storage interface {
	Put(bucket string, object *Object) error
	Get(id string) (*Object, error)
	Delete(id string) error
	Remove(object *Object) error
}
