package stroage

var store Storage

func SetStorage(storage Storage) {
	store = storage
}

func Store() Storage {
	return store
}
