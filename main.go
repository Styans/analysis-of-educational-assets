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

	"fyne.io/fyne/v2/app"
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

	app := app.New()

	i := interface_app.NewInterfaceApp(service, app)
	starter.ShowAndRun(i)

}
