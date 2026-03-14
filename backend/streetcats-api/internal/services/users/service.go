package users

import (
	"context"
	"streetcats-api/configs"
	"streetcats-api/internal/dto"
	"streetcats-api/internal/repositories/users"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ServiceInterfaces interface {
	SetContext(ctx any)
	CreateUser(account dto.AccountDTO) error
	// GetUser(userID string) (*gocloak.User, error)
	// UpdateUser(userID string, user *gocloak.User) error
	// DeleteUser(userID string) error
}

type Service struct {
	log            *zap.Logger
	cfg            *configs.ConfigModel
	kc             *gocloak.GoCloak
	ctx            context.Context
	userRepository users.Repository
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

func NewUsersService(
	log *zap.Logger,
	cfg configs.ConfigModel,
	kc *gocloak.GoCloak,
	//rds *redis.Client,
	userRepository users.Repository,
) *Service {
	return &Service{log: log, cfg: &cfg, kc: kc, userRepository: userRepository}
}

func (s *Service) createKeycloakUser(username, email, password string, language, firstName, lastName *string) error {
	// retrieve admin token
	realm := s.cfg.Keycloak.Realm
	token, err := s.kc.LoginAdmin(s.ctx, s.cfg.Keycloak.User, s.cfg.Keycloak.Passwd, realm)
	if err != nil {
		return err
	}

	// Create user in Keycloak
	attr := map[string][]string{
		"language": {getStringValue(language)},
	}

	user := gocloak.User{
		Username:   gocloak.StringP(username),
		Email:      gocloak.StringP(email),
		Attributes: &attr,
		Enabled:    gocloak.BoolP(true),
	}

	if firstName != nil {
		user.FirstName = firstName
	}
	if lastName != nil {
		user.LastName = lastName
	}

	userID, err := s.kc.CreateUser(s.ctx, token.AccessToken, realm, user)
	if err != nil {
		s.log.Error("Failed to create user in Keycloak", zap.Error(err))
		return err
	}

	err = s.kc.SetPassword(s.ctx, token.AccessToken, userID, realm, password, false)
	if err != nil {
		return err
	}

	s.log.Debug("User created in Keycloak")
	return nil
}

func (s *Service) CreateUser(account dto.AccountDTO) error {

	err := s.createKeycloakUser(account.Username, account.Email, account.Password, account.Language, account.FirstName, account.LastName)
	if err != nil {
		return err
	}

	// Create user in local database
	accountProfile, err := s.userRepository.CreateUser(account.Username, account.Email, account.Language, account.FirstName, account.LastName)
	if err != nil {
		s.log.Error("Failed to create user in local database", zap.Error(err))
		return err
	}

	s.log.Debug("User created in local database", zap.Uint64("accountID", accountProfile.Account.ID))
	s.log.Info("User created in local database")

	return nil
}

func getStringValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return "en"
}
