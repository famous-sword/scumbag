package entity

import (
	"fmt"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/plugger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type DatabasePlugger struct{}

func (d *DatabasePlugger) Plug() (err error) {
	db, err = gorm.Open(resolveDriver(), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return fmt.Errorf("connect database error %s", err)
	}

	err = db.AutoMigrate(&Resource{}, &LocalStorage{})

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

func NewDatabasePlugger() plugger.Plugger {
	return &DatabasePlugger{}
}
