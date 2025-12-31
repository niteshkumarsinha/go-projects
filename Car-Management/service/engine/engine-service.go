package engineService

import (
	"context"

	"github.com/nitesh111sinha/car-management/models"
	"github.com/nitesh111sinha/car-management/store"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}	
}

func (s *EngineService) GetEngineById(ctx context.Context, engineID string) (models.Engine, error) {
	engine, err := s.store.GetEngineById(ctx, engineID)
	if err != nil {
		return models.Engine{}, err
	}
	return engine, nil
}

func (s *EngineService) GetEngines(ctx context.Context) ([]models.Engine, error) {
	engines, err := s.store.GetEngines(ctx)
	if err != nil {
		return nil, err
	}
	return engines, nil
}

func (s *EngineService) UpdateEngine(ctx context.Context, engine models.Engine) (models.Engine, error) {
	updatedEngine, err := s.store.UpdateEngine(ctx, engine)
	if err != nil {
		return models.Engine{}, err
	}
	return updatedEngine, nil
}

func (s *EngineService) DeleteEngine(ctx context.Context, engineID string) error {
	if err := s.store.DeleteEngine(ctx, engineID); err != nil {
		return err
	}
	return nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engine models.Engine) (models.Engine, error) {
	createdEngine, err := s.store.CreateEngine(ctx, engine)
	if err != nil {
		return models.Engine{}, err
	}
	return createdEngine, nil
}
