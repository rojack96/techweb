package sightings

import (
	"context"
	"streetcats-api/configs"
	"streetcats-api/internal/dto"
	"streetcats-api/internal/repositories/sightings"

	"github.com/Nerzal/gocloak/v13"
	"go.uber.org/zap"
)

type ServiceInterfaces interface {
	GetAllSightings(ctx context.Context) ([]dto.SightingDTO, error)
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
		sightingDTOs = append(sightingDTOs, dto.SightingDTO{
			ID:        sighting.ID,
			AnimalID:  sighting.AnimalID,
			BreedID:   sighting.BreedID,
			Position:  [2]float64{sighting.Latitude, sighting.Longitude},
			SpottedAt: sighting.SpottedAt,
		})
	}

	return sightingDTOs, nil
}
