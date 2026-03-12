package controllers

import (
	"github.com/rojack96/jinres"

	"go.uber.org/zap"
)

type Controller struct {
	log    *zap.Logger
	jinres *jinres.Jinres
}

func NewController(log *zap.Logger) *Controller {
	return &Controller{log: log, jinres: jinres.NewJinres()}
}
