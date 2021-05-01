package main

import (
	"context"
	"github.com/famous-sword/scumbag/engine"
	"github.com/famous-sword/scumbag/entity"
	"github.com/famous-sword/scumbag/logger"
	"github.com/famous-sword/scumbag/storage"
	"github.com/famous-sword/scumbag/upload"
	"log"
)

func main() {
	scheduler := engine.NewScheduler()

	// bootstrap components
	scheduler.Register(entity.NewDatabaseBootstrapper())
	scheduler.Register(logger.NewBootstrapper())
	scheduler.Register(storage.NewBootstrapper())

	// register routes
	scheduler.Register(upload.NewUploader())

	if err := scheduler.Bootstrap(); err != nil {
		log.Fatal(err)
	}

	scheduler.Run(context.TODO())
}
