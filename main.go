package main

import (
	"flag"
	"log"
	"myApp/configs"
	starter "myApp/internal/app"
	interface_app "myApp/internal/interface"
	"myApp/internal/repository"
	"myApp/internal/service"
	"myApp/pkg/client/sqlite"
)

func main() {

	configPath := flag.String("config", "config.json", "path to config file")
	flag.Parse()

	cfg, err := configs.GetConfig(*configPath)
	if err != nil {
		log.Println(err)
		return
	}

	db, err := sqlite.OpenDB(cfg.DB.DSN)
	if err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	i := interface_app.NewInterfaceApp(service)
	starter.ShowAndRun(i)

}
