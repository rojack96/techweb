package controllers

import (
	"streetcats-api/internal/services/sightings"

	"github.com/rojack96/jinres"

	"go.uber.org/zap"
)

type Controller struct {
	log             *zap.Logger
	sightingService sightings.ServiceInterfaces
	jinres          *jinres.Jinres
}

func NewController(log *zap.Logger, sightingService sightings.ServiceInterfaces) *Controller {
	return &Controller{log: log, sightingService: sightingService, jinres: jinres.NewJinres()}
}
