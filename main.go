package main

import (
	"fmt"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/engine"
	"github.com/famous-sword/scumbag/entity"
	"github.com/famous-sword/scumbag/stroage"
	"github.com/famous-sword/scumbag/stroage/minio"
	"log"
)

func main() {
	engine.Register(entity.NewDatabasePlugger())

	err := engine.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}

	stroage.SetStorage(minio.NewMinio())

	fmt.Println(config.String("app.name"))
}
