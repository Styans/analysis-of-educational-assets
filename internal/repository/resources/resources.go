package resources

import (
	"context"
	"database/sql"
	"fmt"
	"myApp/internal/models"
	"time"
)

type ResourcesStorage struct {
	db *sql.DB
}

func NewResourceStorage(db *sql.DB) *ResourcesStorage {
	return &ResourcesStorage{db}
}

// func (r *ResourceStorage) CreateResource() {

// }

func (r *ResourcesStorage) GetResourcesByType(t string) ([]models.ResourcesDTO, error) {
	query := `SELECT title, cost, created_at FROM resource WHERE type = $1`
	rows, err := r.db.Query(query, t)
	if err != nil {
		return nil, fmt.Errorf("querying resources by type: %v", err)
	}
	defer rows.Close()

	var resources []models.ResourcesDTO
	for rows.Next() {
		var r models.ResourcesDTO
		err := rows.Scan(&r.Resource, &r.Cost, &r.Date)
		if err != nil {
			return nil, fmt.Errorf("scanning resource: %v", err)
		}
		resources = append(resources, r)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("reading rows: %v", err)
	}

	return resources, nil

}

func (s *ResourcesStorage) CreateResources(r *models.ResourcesDTO) error {
	query := `INSERT INTO resource (title, type, cost, authorname, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	args := []interface{}{r.Resource, r.Category, r.Cost, r.AuthorName, r.Date, r.Date}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
