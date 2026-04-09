package sightings

import (
	"context"
	"streetcats-api/internal/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type sightingsRepositoryImpl struct {
	pg *pgxpool.Pool
}

func (r *sightingsRepositoryImpl) GetAllSightings() ([]entities.Sighting, error) {
	ctx := context.Background()
	query := `
		SELECT id, animal_id, breed_id, latitude, longitude, spotted_at
		FROM sightings.sightings
	`

	rows, err := r.pg.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sightings []entities.Sighting
	for rows.Next() {
		var sighting entities.Sighting
		err := rows.Scan(
			&sighting.ID,
			&sighting.AnimalID,
			&sighting.BreedID,
			&sighting.Latitude,
			&sighting.Longitude,
			&sighting.SpottedAt,
		)
		if err != nil {
			return nil, err
		}
		sightings = append(sightings, sighting)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return sightings, nil
}

func NewRepository(pg *pgxpool.Pool) Repository {
	return &sightingsRepositoryImpl{pg: pg}
}
