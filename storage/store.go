package storage

var adapter Storage

func SetStorage(storage Storage) {
	adapter = storage
}

func Adapter() Storage {
	return adapter
}
