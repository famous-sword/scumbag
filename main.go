package main

import (
	"fmt"
	"github.com/famous-sword/scumbag/api"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/engine"
	"github.com/famous-sword/scumbag/entity"
	"github.com/famous-sword/scumbag/logger"
	"github.com/famous-sword/scumbag/stroage"
	"github.com/famous-sword/scumbag/stroage/local"
	"log"
)

func main() {
	engine.Register(entity.NewDatabasePlugger())
	engine.Register(logger.NewPlugger())

	err := engine.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}

	logger.Writter().Error("hello")

	stroage.SetStorage(local.NewLocal())

	api.Uploader().Run(address())
}

func address() string {
	return fmt.Sprintf("%s:%d", config.String("web.host"), config.Integer("web.port"))
}
