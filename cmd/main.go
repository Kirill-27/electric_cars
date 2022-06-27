package main

import (
	"github.com/Kirill-27/electric_cars"
	"log"
)

func main() {
	srv := new(electric_cars.Server)
	if err := srv.Run("8002"); err != nil {
		log.Fatalf("error when run http server: %s", err.Error())
	}

}
