package service

import (
	"context"
	"github.com/nitesh111sinha/car-management/models"
)

type CarServiceInterface interface {
	GetCarById(ctx context.Context, carID string) (models.Car, error)
	GetCars(ctx context.Context) ([]models.Car, error)
	UpdateCar(ctx context.Context, car models.Car) (models.Car, error)
	DeleteCar(ctx context.Context, carID string) error
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
	CreateCar(ctx context.Context, car models.Car) (models.Car, error)
}

type EngineServiceInterface interface {
	GetEngineById(ctx context.Context, engineID string) (models.Engine, error)
	GetEngines(ctx context.Context) ([]models.Engine, error)
	UpdateEngine(ctx context.Context, engine models.Engine) (models.Engine, error)
	DeleteEngine(ctx context.Context, engineID string) error
	CreateEngine(ctx context.Context, engine models.Engine) (models.Engine, error)
}	