package engine

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/nitesh111sinha/car-management/models"
)

type EngineStore struct {
	db *sql.DB
}

func NewEngineStore(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (s EngineStore) CreateEngine(ctx context.Context, engine models.Engine) (models.Engine, error) {
	var createdEngine models.Engine

	// Begin Transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createdEngine, err
	}

	engineId := uuid.New()

	newEngine := models.Engine{
		EngineID:      engineId,
		Displacement:  engine.Displacement,
		NoOfCylinders: engine.NoOfCylinders,
		CarRange:      engine.CarRange,
	}

	query := `INSERT INTO engines (id, displacement, no_of_cylinders, car_range) VALUES ($1, $2, $3, $4) RETURNING id, displacement, no_of_cylinders, car_range`

	err = tx.QueryRowContext(ctx, query, newEngine.EngineID, newEngine.Displacement, newEngine.NoOfCylinders, newEngine.CarRange).Scan(
		&createdEngine.EngineID,
		&createdEngine.Displacement,
		&createdEngine.NoOfCylinders,
		&createdEngine.CarRange)

	if err != nil {
		tx.Rollback()
		return createdEngine, err
	}

	// Commit Transaction
	if err = tx.Commit(); err != nil {
		return createdEngine, err
	}

	return createdEngine, nil
}

func (s EngineStore) UpdateEngine(ctx context.Context, engineId string, engine *models.Engine) (models.Engine, error) {
	var updatedEngine models.Engine

	// Begin Transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return updatedEngine, err
	}

	// Update Engine
	query := `UPDATE engines SET displacement=$2, no_of_cylinders=$3, car_range=$4 WHERE id=$1 RETURNING id, displacement, no_of_cylinders, car_range`

	err = tx.QueryRowContext(ctx, query, engineId, engine.Displacement, engine.NoOfCylinders, engine.CarRange).Scan(
		&updatedEngine.EngineID,
		&updatedEngine.Displacement,
		&updatedEngine.NoOfCylinders,
		&updatedEngine.CarRange)

	if err != nil {
		tx.Rollback()
		return updatedEngine, err
	}

	// Commit Transaction
	if err = tx.Commit(); err != nil {
		return updatedEngine, err
	}

	return updatedEngine, nil
}

func (s EngineStore) GetEngineById(ctx context.Context, engineId string) (models.Engine, error) {
	var engine models.Engine

	query := `SELECT id, displacement, no_of_cylinders, car_range FROM engines WHERE id=$1`

	err := s.db.QueryRowContext(ctx, query, engineId).Scan(
		&engine.EngineID,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange)
	if err != nil {
		return engine, err
	}

	return engine, nil
}

func (s EngineStore) DeleteEngine(ctx context.Context, engineId string) error {
	// Begin Transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Delete Engine
	query := `DELETE FROM engines WHERE id=$1`

	row := tx.QueryRowContext(ctx, query, engineId)
	if row.Scan() != nil {
		tx.Rollback()
		return err
	}

	// Commit Transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s EngineStore) GetEngines(ctx context.Context) ([]models.Engine, error) {
	var engines []models.Engine

	query := `SELECT id, displacement, no_of_cylinders, car_range FROM engines`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return engines, err
	}
	defer rows.Close()

	for rows.Next() {
		var engine models.Engine
		err := rows.Scan(
			&engine.EngineID,
			&engine.Displacement,
			&engine.NoOfCylinders,
			&engine.CarRange)
		if err != nil {
			return engines, err
		}
		engines = append(engines, engine)
	}

	if err := rows.Err(); err != nil {
		return engines, err
	}

	return engines, nil
}
