package sightings

import (
	"streetcats-api/internal/entities"
)

type Repository interface {
	GetAllSightings() ([]entities.AnimalEntitiesView, error)
	BreedsLookup(animalId uint64) (*entities.Breed, error)
	CreateSighting(sighting entities.AnimalEntities) (uint64, error)
}
