package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nitesh111sinha/car-management/models"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
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
	var createdCar models.Car
	var engineId uuid.UUID

	engineRow := s.db.QueryRowContext(ctx, `SELECT id FROM engine WHERE id=$1`, car.Engine.EngineID)

	if engineRow.Scan(&engineId) != nil {
		return createdCar, errors.New("engine id is required and must be a valid uuid")
	}

	carId := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt

	newCar := models.Car{
		ID:        carId,
		Name:      car.Name,
		Year:      car.Year,
		Brand:     car.Brand,
		FuelType:  car.FuelType,
		Engine:    car.Engine,
		Price:     car.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	// Begin Transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createdCar, err
	}

	// Insert Car
	query := `INSERT INTO cars (id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`

	err = tx.QueryRowContext(ctx, query,
		newCar.ID,
		newCar.Name,
		newCar.Year,
		newCar.Brand,
		newCar.FuelType,
		newCar.Engine.EngineID,
		newCar.Price,
		newCar.CreatedAt,
		newCar.UpdatedAt).Scan(
		&createdCar.ID,
		&createdCar.Name,
		&createdCar.Year,
		&createdCar.Brand,
		&createdCar.FuelType,
		&createdCar.Engine.EngineID,
		&createdCar.Price,
		&createdCar.CreatedAt,
		&createdCar.UpdatedAt)

	if err != nil {

		tx.Rollback()
		return createdCar, err
	}

	// Commit Transaction
	if err = tx.Commit(); err != nil {
		return createdCar, err
	}

	return createdCar, nil
}

func (s Store) UpdateCar(ctx context.Context, car models.Car) (models.Car, error) {
	var updatedCar models.Car
	var engineId uuid.UUID

	engineRow := s.db.QueryRowContext(ctx, "SELECT id FROM engine WHERE id=$1", car.Engine.EngineID)

	if engineRow.Scan(&engineId) != nil {
		return updatedCar, errors.New("engine id is required and must be a valid uuid")
	}

	// Begin Transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return updatedCar, err
	}

	// Update Car
	query := `UPDATE cars SET name=$2, year=$3, brand=$4, fuel_type=$5, engine_id=$6, price=$7, updated_at=$8 WHERE id=$1 RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`

	err = tx.QueryRowContext(ctx, query,
		car.ID,
		car.Name,
		car.Year,
		car.Brand,
		car.FuelType,
		car.Engine.EngineID,
		car.Price,
		car.UpdatedAt).Scan(
		&updatedCar.ID,
		&updatedCar.Name,
		&updatedCar.Year,
		&updatedCar.Brand,
		&updatedCar.FuelType,
		&updatedCar.Engine.EngineID,
		&updatedCar.Price,
		&updatedCar.CreatedAt,
		&updatedCar.UpdatedAt)

	if err != nil {

		tx.Rollback()
		return updatedCar, err
	}

	// Commit Transaction
	if err = tx.Commit(); err != nil {
		return updatedCar, err
	}

	return updatedCar, nil
}

func (s Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	var deletedCar models.Car
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return deletedCar, err
	}

	// Delete Car
	query := `DELETE FROM cars WHERE id=$1 RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`

	err = tx.QueryRowContext(ctx, query, id).Scan(
		&deletedCar.ID,
		&deletedCar.Name,
		&deletedCar.Year,
		&deletedCar.Brand,
		&deletedCar.FuelType,
		&deletedCar.Engine.EngineID,
		&deletedCar.Price,
		&deletedCar.CreatedAt,
		&deletedCar.UpdatedAt)

	if err != nil {
		tx.Rollback()
		return deletedCar, err
	}

	// Commit Transaction
	if err = tx.Commit(); err != nil {
		return deletedCar, err
	}

	return deletedCar, nil
}

func (s Store) GetCars(ctx context.Context) ([]models.Car, error) {
	var cars []models.Car
	query := `SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at FROM cars`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return cars, err
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car
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
			return cars, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}
