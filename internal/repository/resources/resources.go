package resources

import (
	"database/sql"
)

type ResourcesStorage struct {
	db *sql.DB
}

func NewResourceStorage(db *sql.DB) *ResourcesStorage {
	return &ResourcesStorage{db}
}

// func (r *ResourceStorage) CreateResource() {

// }

func (r *ResourcesStorage) GetResourcesTypes() ([]string, error) {
	query := `SELECT * FROM resources `
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var resources []string
	for rows.Next() {
		err := rows.Scan(&resources)
		if err != nil {
			return nil, err
		}
	}
	return resources, nil
}
