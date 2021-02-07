package main

import (
	"context"
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
	scheduler := engine.NewScheduler()

	scheduler.Register(entity.NewDatabasePlugger())
	scheduler.Register(logger.NewPlugger())

	err := scheduler.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}

	stroage.SetStorage(local.NewLocal())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go api.Uploader().Run(address())

	scheduler.Run(ctx)
}

func address() string {
	return fmt.Sprintf("%s:%d", config.String("web.host"), config.Integer("web.port"))
}
