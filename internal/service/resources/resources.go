package resources

import "myApp/internal/models"

type ResourcesService struct {
	repo models.ResourcesRepo
}

func NewResourcesService(repo models.ResourcesRepo) *ResourcesService {
	return &ResourcesService{repo}
}

func (s *ResourcesService) GetResourcesTypes() ([]string, error) {
	return s.repo.GetResourcesTypes()
}
