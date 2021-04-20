package storage

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/foundation"
	"github.com/famous-sword/scumbag/storage/driver"
)

var _driver foundation.StorageDriver

type Bootstrapper struct{}

func (_ *Bootstrapper) Bootstrap() (err error) {
	var s foundation.StorageDriver
	switch config.String("storage.driver") {
	case "minio":
		s, err = driver.NewMinio()
	case "ceph":
		s, err = driver.NewCeph()
	default:
		s = driver.NewLocal()
	}

	if err != nil {
		return err
	}

	SetDriver(s)

	return err
}

func NewBootstrapper() foundation.Bootable {
	return &Bootstrapper{}
}

// SetDriver for change driver runtime
func SetDriver(storage foundation.StorageDriver) {
	_driver = storage
}

func Driver() foundation.StorageDriver {
	return _driver
}
