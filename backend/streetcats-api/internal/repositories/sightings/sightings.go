package sightings

import (
	"streetcats-api/internal/entities"
)

type Repository interface {
	GetAllSightings() ([]entities.AnimalEntitiesView, error)
}
