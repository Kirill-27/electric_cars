package main

import (
	"github.com/Kirill-27/electric_cars"
	"github.com/Kirill-27/electric_cars/pkg/handler"
	"github.com/Kirill-27/electric_cars/pkg/repository"
	"github.com/Kirill-27/electric_cars/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := new(handler.Handler)
	srv := new(electric_cars.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("error when run http server: %s", err.Error())
	}

}

func initConfig() error {
	//viper.AddConfigPath()
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
