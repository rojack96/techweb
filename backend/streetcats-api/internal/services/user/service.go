package auth

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ServiceInterfaces interface {
	CreateUser() (string, error)
	// GetUser(userID string) (*gocloak.User, error)
	// UpdateUser(userID string, user *gocloak.User) error
	// DeleteUser(userID string) error
}

type Service struct {
	log *zap.Logger
	kc  *gocloak.GoCloak
	ctx context.Context
	//txManager        gorm.TransactionManager
}

func (s *Service) SetContext(ctx any) {
	switch v := ctx.(type) {
	case *gin.Context:
		s.ctx = v.Request.Context()
	case context.Context:
		s.ctx = v
	default:
		s.ctx = context.Background()
	}
}

func NewNotificationService(
	log *zap.Logger,
	kc *gocloak.GoCloak,
	//rds *redis.Client,
	//txManager gorm.TransactionManager,
) *Service {
	return &Service{log: log, kc: kc}
}
