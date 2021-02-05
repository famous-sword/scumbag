package main

import (
	"fmt"
	"github.com/famous-sword/scumbag/config"
	"github.com/famous-sword/scumbag/engine"
	"github.com/famous-sword/scumbag/entity"
	"log"
)

func main() {
	engine.Register(entity.NewDatabasePlugger())

	err := engine.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.String("app.name"))
}
