package students

import "myApp/internal/models"

type StudentsService struct {
	repo models.StudentsRepo
}

func NewStudentService(repo models.StudentsRepo) *StudentsService {
	return &StudentsService{repo}
}
