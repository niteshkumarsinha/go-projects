package store

import (
	"context"
	"database/sql"
	"github.com/nitesh111sinha/car-management/models"
)

type Store struct {
	db *sql.DB
}	

func new(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	var car models.Car
	query := `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cylinders, e.car_range	 FROM cars c LEFT JOIN engines e ON c.engine_id = e.engine_id WHERE c.id=$1`

	row := s.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&car.ID, 
					&car.Name, 
					&car.Year, 
					&car.Brand, 
					&car.FuelType, 
					&car.Engine.EngineID, 
					&car.Price, 
					&car.CreatedAt, 
					&car.UpdatedAt, 
					&car.Engine.EngineID, 
					&car.Engine.Displacement, 
					&car.Engine.NoOfCylinders, 
					&car.Engine.CarRange)

	if err != nil {
		if err == sql.ErrNoRows {
			return car, nil
		}
		return car, err
	}
	return car, nil
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars []models.Car
	var query string
	if isEngine {
		query = `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cylinders, e.car_range FROM cars c LEFT JOIN engines e ON c.engine_id = e.engine_id WHERE c.brand=$1`
	} else {
		query = `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at FROM cars c WHERE c.brand=$1`
	}

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return cars, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car
		if isEngine {
			var engine models.Engine
			err := rows.Scan(&car.ID, 
					&car.Name, 
					&car.Year, 
					&car.Brand, 
					&car.FuelType, 
					&car.Engine.EngineID, 
					&car.Price, 
					&car.CreatedAt, 
					&car.UpdatedAt, 
					&engine.EngineID, 
					&engine.Displacement, 
					&engine.NoOfCylinders, 
					&engine.CarRange)
			if err != nil {
				return nil, err
			}
			car.Engine = engine
			cars = append(cars, car)
		} else {
			err := rows.Scan(&car.ID, 
							&car.Name, 
							&car.Year, 
							&car.Brand, 
							&car.FuelType, 
							&car.Engine.EngineID, 
							&car.Price, 
							&car.CreatedAt, 
							&car.UpdatedAt)
			if err != nil {
				return nil, err
			}
			cars = append(cars, car)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}	

	return cars, nil
}

func (s Store) CreateCar(ctx context.Context, car models.Car) (models.Car, error) {
	return models.Car{}, nil
}

func (s Store) UpdateCar(ctx context.Context, car models.Car) (models.Car, error) {
	return models.Car{}, nil
}

func (s Store) DeleteCar(ctx context.Context, id string) error {
	return nil
}	

func (s Store) GetCarByYear(ctx context.Context, year string) ([]models.Car, error) {
	return []models.Car{}, nil
}	



