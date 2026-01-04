package engineService

import (
	"context"

	"github.com/nitesh111sinha/car-management/models"
	"github.com/nitesh111sinha/car-management/store"
	"go.opentelemetry.io/otel"
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
	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "GetEngineById-Service")
	defer span.End()
	engine, err := s.store.GetEngineById(ctx, engineID)
	if err != nil {
		return models.Engine{}, err
	}
	return engine, nil
}

func (s *EngineService) GetEngines(ctx context.Context) ([]models.Engine, error) {
	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "GetEngines-Service")
	defer span.End()
	engines, err := s.store.GetEngines(ctx)
	if err != nil {
		return nil, err
	}
	return engines, nil
}

func (s *EngineService) UpdateEngine(ctx context.Context, engineID string, engine models.Engine) (models.Engine, error) {
	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "UpdateEngine-Service")
	defer span.End()
	updatedEngine, err := s.store.UpdateEngine(ctx, engineID, engine)
	if err != nil {
		return models.Engine{}, err
	}
	return updatedEngine, nil
}

func (s *EngineService) DeleteEngine(ctx context.Context, engineID string) error {
	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "DeleteEngine-Service")
	defer span.End()
	if err := s.store.DeleteEngine(ctx, engineID); err != nil {
		return err
	}
	return nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engine models.Engine) (models.Engine, error) {
	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "CreateEngine-Service")
	defer span.End()
	createdEngine, err := s.store.CreateEngine(ctx, engine)
	if err != nil {
		return models.Engine{}, err
	}
	return createdEngine, nil
}
