package main

import (
	"context"
	"github.com/famous-sword/scumbag/engine"
	"github.com/famous-sword/scumbag/entity"
	"github.com/famous-sword/scumbag/logger"
	"github.com/famous-sword/scumbag/storage"
	"github.com/famous-sword/scumbag/storage/local"
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

	storage.SetStorage(local.NewLocal())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	scheduler.Run(ctx)
}
