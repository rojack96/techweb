package auth

import (
	"streetcats-api/internal/services/keycloak"
	"streetcats-api/internal/services/session"

	"github.com/rojack96/jinres"
	"go.uber.org/zap"
)

type Controller struct {
	log            *zap.Logger
	kcService      keycloak.ServiceInterfaces
	sessionService session.ServiceInterfaces
	jinres         *jinres.Jinres
}

func NewAuthController(log *zap.Logger, kcService keycloak.ServiceInterfaces, sessionService session.ServiceInterfaces) *Controller {
	return &Controller{log: log, kcService: kcService, sessionService: sessionService, jinres: jinres.NewJinres()}
}
