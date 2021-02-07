package config

import (
	"fmt"
	"github.com/famous-sword/scumbag/plugger"
	"github.com/spf13/viper"
)

type Plugger struct{}

func (c *Plugger) Plug() (err error) {
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

func NewPlugger() plugger.Plugger {
	return &Plugger{}
}
