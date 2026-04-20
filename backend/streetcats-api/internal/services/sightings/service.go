package sightings

import (
	"context"
	"fmt"
	"streetcats-api/configs"
	"streetcats-api/internal/dto"
	"streetcats-api/internal/entities"
	"streetcats-api/internal/repositories/sightings"
	"strings"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"go.uber.org/zap"
)

type ServiceInterfaces interface {
	GetAllSightings(ctx context.Context) ([]dto.SightingDTO, error)
	BreedsLookup(animalId uint64) ([]dto.BreedDTO, error)
	CreateSighting(sighting dto.CreateSightingDTO) (uint64, error)
}

type Service struct {
	log                *zap.Logger
	cfg                *configs.ConfigModel
	kc                 *gocloak.GoCloak
	sightingRepository sightings.Repository
}

func NewService(
	log *zap.Logger,
	cfg configs.ConfigModel,
	kc *gocloak.GoCloak,
	sightingRepository sightings.Repository,
) *Service {
	return &Service{log: log, cfg: &cfg, kc: kc, sightingRepository: sightingRepository}
}

func (s *Service) GetAllSightings(ctx context.Context) ([]dto.SightingDTO, error) {
	sightings, err := s.sightingRepository.GetAllSightings()
	if err != nil {
		return nil, err
	}

	var sightingDTOs []dto.SightingDTO
	for _, sighting := range sightings {
		if sighting.Breed == nil {
			unknown := "unknown"
			sighting.Breed = &unknown
		}

		breed := strings.ReplaceAll(*sighting.Breed, " ", "-")
		// TODO verificarfe che serva
		id := fmt.Sprintf("%s-%d", strings.ToLower(breed), sighting.ID)

		sightingDTOs = append(sightingDTOs, dto.SightingDTO{
			ID:          id,
			SightingID:  sighting.ID,
			Breed:       sighting.Breed,
			Position:    [2]float64{sighting.Latitude, sighting.Longitude},
			Title:       sighting.Title,
			Description: sighting.Description,
			SpottedAt:   sighting.SpottedAt,
		})
	}

	return sightingDTOs, nil
}

func (s *Service) BreedsLookup(animalId uint64) ([]dto.BreedDTO, error) {

	breeds, err := s.sightingRepository.BreedsLookup(animalId)
	if err != nil {
		return nil, err
	}

	result := make([]dto.BreedDTO, len(breeds))
	for i, breed := range breeds {
		result[i] = dto.BreedDTO{
			ID:   breed.ID,
			Name: breed.Name,
		}
	}

	return result, nil
}

func (s *Service) CreateSighting(sighting dto.CreateSightingDTO) (uint64, error) {
	now := time.Now().Unix()

	newSighting := entities.AnimalEntities{
		AnimalID:    sighting.AnimalID,
		BreedID:     &sighting.BreedID,
		Latitude:    sighting.Position[0],
		Longitude:   sighting.Position[1],
		Title:       sighting.Title,
		Description: sighting.Description,
		SpottedAt:   sighting.SpottedAt,
		CreatedAt:   now,
	}

	return s.sightingRepository.CreateSighting(newSighting)
}
