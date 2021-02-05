package model

import (
	"fmt"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/engine"
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

	return nil
}

func resolveDriver() gorm.Dialector {
	var driver gorm.Dialector
	dsn := config.String("database.dsn")

	switch config.String("database.driver") {
	case "mysql":
		driver = mysql.Open(dsn)
	case "sqlite":
		fallthrough
	default:
		driver = sqlite.Open(dsn)
	}

	return driver
}

func NewDatabasePlugger() engine.Plugger {
	return &DatabasePlugger{}
}
