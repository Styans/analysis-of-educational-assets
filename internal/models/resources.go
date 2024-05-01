package models

import "time"

type ResourcesDTO struct {
	Resource   string
	Cost       int
	Category   string
	AuthorName string
	Date       time.Time
}

type ResourcesService interface {
	GetResourcesByType(t string) ([][]string, error)
	CreateResources(r *ResourcesDTO) error
}

type ResourcesRepo interface {
	GetResourcesByType(t string) ([]ResourcesDTO, error)
	CreateResources(r *ResourcesDTO) error
}
