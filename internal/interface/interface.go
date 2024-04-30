package interface_app

import (
	"myApp/internal/service"
)

type Interface_app struct {
	service *service.Service
}

func NewInterfaceApp(service *service.Service) *Interface_app {
	return &Interface_app{service: service}
}
