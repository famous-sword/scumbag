package config

import (
	"fmt"
	"github.com/famous-sword/scumbag/foundation"
	"github.com/spf13/viper"
)

type Bootstrapper struct{}

func (_ *Bootstrapper) Bootstrap() (err error) {
	viper.SetConfigName("app.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./var/")
	viper.AddConfigPath("/etc/scumbag/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("load config error: %s", err)
	}

	return nil
}

func NewBootstrapper() foundation.Bootable {
	return &Bootstrapper{}
}
