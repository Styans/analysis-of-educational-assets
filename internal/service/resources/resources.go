package resources

import (
	"myApp/internal/models"
	"strconv"
)

type ResourcesService struct {
	repo models.ResourcesRepo
}

func NewResourcesService(repo models.ResourcesRepo) *ResourcesService {
	return &ResourcesService{repo}
}

func (s *ResourcesService) GetResourcesByType(t string) ([][]string, error) {
	res, err := s.repo.GetResourcesByType(t)
	if err != nil {
		return nil, err
	}
	data := ConvertResourcesToSliceOfStringSlices(res)
	return data, nil
}

func (s *ResourcesService) CreateResources(r *models.ResourcesDTO) error {
	return s.repo.CreateResources(r)
}

func ConvertResourcesToSliceOfStringSlices(resources []models.ResourcesDTO) [][]string {
	var result [][]string

	for _, r := range resources {
		var row []string
		row = append(row, r.Resource)
		row = append(row, strconv.Itoa(r.Cost))
		row = append(row, r.Date.Format("2006-01-02 15:04:05")) // Форматируем время в строку
		result = append(result, row)
	}
	return result
}
