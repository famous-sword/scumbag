package entity

import (
	"fmt"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/foundation"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Bootstrapper struct{}

func (d *Bootstrapper) Bootstrap() (err error) {
	db, err = gorm.Open(resolveDriver(), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return fmt.Errorf("connect database error %s", err)
	}

	err = db.AutoMigrate(&Resource{})

	if err != nil {
		return fmt.Errorf("migration error %s", err)
	}

	return nil
}

func resolveDriver() (driver gorm.Dialector) {
	dsn := config.String("database.dsn")

	switch config.String("database.driver") {
	case "mysql":
		driver = mysql.Open(dsn)
	case "sqlite", "sqlite3":
		fallthrough
	default:
		driver = sqlite.Open(dsn)
	}

	return driver
}

func NewDatabaseBootstrapper() foundation.Bootable {
	return &Bootstrapper{}
}
