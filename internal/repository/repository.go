package repository

import (
	"database/sql"
	"myApp/internal/models"
	"myApp/internal/repository/resources"
	"myApp/internal/repository/students"
)

type Repository struct {
	StudentsRepo  models.StudentsRepo
	ResourcesRepo models.ResourcesRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		StudentsRepo:  students.NewStudentsStorage(db),
		ResourcesRepo: resources.NewResourceStorage(db),
	}
}
