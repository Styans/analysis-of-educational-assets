package models

type Resources struct {
}

type ResourcesService interface {
	GetResourcesTypes() ([]string, error)
}

type ResourcesRepo interface {
	GetResourcesTypes() ([]string, error)
}
