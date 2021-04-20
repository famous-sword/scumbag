package main

import (
	"context"
	"github.com/famous-sword/scumbag/api"
	"github.com/famous-sword/scumbag/engine"
	"github.com/famous-sword/scumbag/entity"
	"github.com/famous-sword/scumbag/logger"
	"github.com/famous-sword/scumbag/storage"
	"log"
)

func main() {
	scheduler := engine.NewScheduler()

	scheduler.Register(entity.NewDatabaseBootstrapper())
	scheduler.Register(logger.NewBootstrapper())
	scheduler.Register(storage.NewBootstrapper())
	scheduler.Register(api.NewUploader())

	if err := scheduler.Bootstrap(); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	scheduler.Run(ctx)
}
