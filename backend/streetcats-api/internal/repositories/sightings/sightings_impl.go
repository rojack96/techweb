package sightings

import (
	"context"
	"fmt"
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

func (r *sightingsRepositoryImpl) BreedsLookup(animalId uint64) ([]entities.Breed, error) {
	ctx := context.Background()
	query := `
		SELECT id, name
		FROM sightings.breeds
		WHERE animal_id = $1
	`

	var breeds []entities.Breed
	rows, err := r.pg.Query(ctx, query, animalId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var breed entities.Breed
		err := rows.Scan(&breed.ID, &breed.Name)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, breed)
	}

	if len(breeds) == 0 {
		return nil, fmt.Errorf("no breeds found for animal ID: %d", animalId)
	}

	return breeds, nil
}

func (r *sightingsRepositoryImpl) CreateSighting(sighting entities.AnimalEntities) (uint64, error) {
	ctx := context.Background()
	query := `
		INSERT INTO sightings.animal_entities (animal_id, breed_id, latitude, longitude, title, description, spotted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var id uint64
	err := r.pg.QueryRow(ctx, query,
		sighting.AnimalID,
		sighting.BreedID,
		sighting.Latitude,
		sighting.Longitude,
		sighting.Title,
		sighting.Description,
		sighting.SpottedAt,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func NewRepository(pg *pgxpool.Pool) Repository {
	return &sightingsRepositoryImpl{pg: pg}
}
