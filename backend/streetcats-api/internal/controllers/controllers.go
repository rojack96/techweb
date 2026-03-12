package controllers

import (
	"streetcats-api/internal/services/users"

	"github.com/rojack96/jinres"

	"go.uber.org/zap"
)

type Controller struct {
	log          *zap.Logger
	usersService users.ServiceInterfaces
	jinres       *jinres.Jinres
}

func NewController(log *zap.Logger, usersService users.ServiceInterfaces) *Controller {
	return &Controller{log: log, usersService: usersService, jinres: jinres.NewJinres()}
}
