package main

import (
	"github.com/famous-sword/scumbag/engine"
	"github.com/famous-sword/scumbag/model"
	"log"
)

func main() {
	engine.Register(model.NewDatabasePlugger())

	err := engine.Bootstrap()

	if err != nil {
		log.Fatal(err)
	}
}
