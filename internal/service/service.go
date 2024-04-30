package service

import (
	"myApp/internal/models"
	"myApp/internal/repository"
	"myApp/internal/service/resources"
	"myApp/internal/service/students"
)

type Service struct {
	StudentsService  models.StudentsService
	ResourcesService models.ResourcesService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		StudentsService:  students.NewStudentService(repo.StudentsRepo),
		ResourcesService: resources.NewResourcesService(repo.ResourcesRepo),
	}
}
