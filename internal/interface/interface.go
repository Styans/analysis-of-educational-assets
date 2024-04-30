package interface_app

import (
	"myApp/internal/service"

	"fyne.io/fyne/v2"
)

type Interface_app struct {
	Window  fyne.Window
	service *service.Service
	App     fyne.App
}

func NewInterfaceApp(service *service.Service, a fyne.App) *Interface_app {
	return &Interface_app{
		Window:  a.NewWindow("Анализатор"),
		service: service,
		App:     a,
	}
}
