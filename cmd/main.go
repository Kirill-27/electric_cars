package main

import (
	"github.com/Kirill-27/electric_cars"
	"github.com/Kirill-27/electric_cars/pkg/handler"
	"log"
)

func main() {
	srv := new(electric_cars.Server)
	handler := new(handler.Handler)
	if err := srv.Run("8002", handler.InitRouters()); err != nil {
		log.Fatalf("error when run http server: %s", err.Error())
	}

}
