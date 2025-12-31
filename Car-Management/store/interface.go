package store

import (
	"context"
	"github.com/nitesh111sinha/car-management/models"
)

type CarStoreInterface interface {
	CreateCar(ctx context.Context, car models.Car) (models.Car, error)
	GetCarById(ctx context.Context, carID string) (models.Car, error)
	GetCars(ctx context.Context) ([]models.Car, error)
	UpdateCar(ctx context.Context, car models.Car) (models.Car, error)
	DeleteCar(ctx context.Context, carID string) error
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
}

type EngineStoreInterface interface {
	CreateEngine(ctx context.Context, engine models.Engine) (models.Engine, error)
	GetEngineById(ctx context.Context, engineID string) (models.Engine, error)
	GetEngines(ctx context.Context) ([]models.Engine, error)
	UpdateEngine(ctx context.Context, engine models.Engine) (models.Engine, error)
	DeleteEngine(ctx context.Context, engineID string) error
}	