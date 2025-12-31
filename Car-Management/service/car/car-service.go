package carService

import (
	"context"

	"github.com/nitesh111sinha/car-management/models"
	"github.com/nitesh111sinha/car-management/store"
)

type CarService struct {
	store store.CarStoreInterface
}

func NewCarService(store store.CarStoreInterface) *CarService {
	return &CarService{
		store: store,
	}
}

func (s *CarService) GetCarById(ctx context.Context, carID string) (models.Car, error) {
	car, err := s.store.GetCarById(ctx, carID)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (s *CarService) GetCars(ctx context.Context) ([]models.Car, error) {
	cars, err := s.store.GetCars(ctx)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *CarService) UpdateCar(ctx context.Context, car models.Car) (models.Car, error) {
	updatedCar, err := s.store.UpdateCar(ctx, car)
	if err != nil {
		return models.Car{}, err
	}
	return updatedCar, nil
}

func (s *CarService) DeleteCar(ctx context.Context, carID string) error {
	if err := s.store.DeleteCar(ctx, carID); err != nil {
		return err
	}
	return nil
}

func (s *CarService) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	cars, err := s.store.GetCarByBrand(ctx, brand, isEngine)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *CarService) CreateCar(ctx context.Context, car models.Car) (models.Car, error) {
	createdCar, err := s.store.CreateCar(ctx, car)
	if err != nil {
		return models.Car{}, err
	}
	return createdCar, nil
}

