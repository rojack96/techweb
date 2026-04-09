package sightings

import (
	"context"
	"streetcats-api/internal/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type sightingsRepositoryImpl struct {
	pg *pgxpool.Pool
}

func (r *sightingsRepositoryImpl) GetAllSightings() ([]entities.AnimalEntitiesView, error) {
	ctx := context.Background()
	query := `
		SELECT ae.id, ae.animal_id, b.name as breed, ae.latitude, ae.longitude, ae.title, ae.description, ae.spotted_at, ae.created_at
		FROM sightings.animal_entities ae
			LEFT JOIN sightings.breeds b ON ae.breed_id = b.id
	`

	rows, err := r.pg.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sightings []entities.AnimalEntitiesView
	for rows.Next() {
		var sighting entities.AnimalEntitiesView
		err := rows.Scan(
			&sighting.ID,
			&sighting.AnimalID,
			&sighting.Breed,
			&sighting.Latitude,
			&sighting.Longitude,
			&sighting.Title,
			&sighting.Description,
			&sighting.SpottedAt,
			&sighting.CreatedAt,
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
