package store

import (
	"context"
	"database/sql"

	"github.com/nitesh111sinha/car-management/models"
)


type EngineStore struct {
	db *sql.DB
}

func new(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (s EngineStore) CreateEngine(ctx context.Context, engine models.Engine) (models.Engine, error) {
	return models.Engine{}, nil
}

func (s EngineStore) UpdateEngine(ctx context.Context, engineId string, engineRequest *models.EngineRequest) (models.Engine, error) {
	return models.Engine{}, nil
}

func (s EngineStore) GetEngineById(ctx context.Context, engineId string) (models.Engine, error) {
	return models.Engine{}, nil
}	

func (s EngineStore) DeleteEngine(ctx context.Context, engineId string) error {
	return nil
}

func (s EngineStore) GetEngines(ctx context.Context) ([]models.Engine, error) {
	return []models.Engine{}, nil
}

