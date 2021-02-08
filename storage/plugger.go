package storage

import (
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/plugger"
	"github.com/famous-sword/scumbag/storage/driver"
)

var _driver driver.StorageDriver

type Plugger struct{}

func (p *Plugger) Plug() (err error) {
	var s driver.StorageDriver
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

func NewPlugger() plugger.Plugger {
	return &Plugger{}
}

// for change driver runtime
func SetDriver(storage driver.StorageDriver) {
	_driver = storage
}

func Driver() driver.StorageDriver {
	return _driver
}
