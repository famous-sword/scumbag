package storage

var driver Storage

func SetDriver(storage Storage) {
	driver = storage
}

func Driver() Storage {
	return driver
}
