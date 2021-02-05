package stroage

type Storage interface {
	Put(object *Object) error
	Get(id string) *Object
	Delete(id string) error
	Remove(object *Object) error
}
