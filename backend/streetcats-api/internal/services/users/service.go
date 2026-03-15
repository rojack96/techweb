package users

import (
	"context"
	"fmt"
	"streetcats-api/configs"
	"streetcats-api/internal/dto"
	"streetcats-api/internal/entities"
	"streetcats-api/internal/repositories/users"

	"github.com/Nerzal/gocloak/v13"
	"go.uber.org/zap"
)

type ServiceInterfaces interface {
	CreateUser(ctx context.Context, account dto.AccountDTO) error
	ResetPassword(ctx context.Context, identifier, newPassword string) error
	GetUsernameByEmail(ctx context.Context, email string) (string, error)
	// GetUser(userID string) (*gocloak.User, error)
	// UpdateUser(userID string, user *gocloak.User) error
	// DeleteUser(userID string) error
}

type Service struct {
	log            *zap.Logger
	cfg            *configs.ConfigModel
	kc             *gocloak.GoCloak
	userRepository users.Repository
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

func (s *Service) createKeycloakUser(ctx context.Context, account dto.AccountDTO) error {
	// retrieve admin token
	realm := s.cfg.Keycloak.Realm
	token, err := s.kc.LoginAdmin(ctx, s.cfg.Keycloak.User, s.cfg.Keycloak.Passwd, realm)
	if err != nil {
		return err
	}

	// Create user in Keycloak
	attr := map[string][]string{
		"language": {getStringValue(account.Language)},
	}

	user := gocloak.User{
		Username:   gocloak.StringP(account.Username),
		Email:      gocloak.StringP(account.Email),
		FirstName:  account.FirstName,
		LastName:   account.LastName,
		Attributes: &attr,
		Enabled:    gocloak.BoolP(true),
	}

	userID, err := s.kc.CreateUser(ctx, token.AccessToken, realm, user)
	if err != nil {
		s.log.Error("Failed to create user in Keycloak", zap.Error(err))
		return err
	}

	err = s.kc.SetPassword(ctx, token.AccessToken, userID, realm, account.Password, false)
	if err != nil {
		return err
	}

	s.log.Debug("User created in Keycloak")
	return nil
}

func (s *Service) CreateUser(ctx context.Context, account dto.AccountDTO) error {

	err := s.createKeycloakUser(ctx, account)
	if err != nil {
		return err
	}

	accountEntity := entities.Account{
		Username: account.Username,
		Email:    account.Email,
		Language: account.Language,
	}

	profileEntity := entities.Profile{
		FirstName: account.FirstName,
		LastName:  account.LastName,
	}

	// Create user in local database
	accountProfile, err := s.userRepository.CreateUser(ctx, accountEntity, profileEntity)
	if err != nil {
		s.log.Error("Failed to create user in local database", zap.Error(err))
		return err
	}

	s.log.Debug("User created in local database", zap.Uint64("accountID", accountProfile.Account.ID))
	s.log.Info("User created in local database")

	return nil
}

func (s *Service) ResetPassword(ctx context.Context, identifier, newPassword string) error {
	// retrieve admin token
	realm := s.cfg.Keycloak.Realm
	token, err := s.kc.LoginAdmin(ctx, s.cfg.Keycloak.User, s.cfg.Keycloak.Passwd, realm)
	if err != nil {
		s.log.Error("Failed to login admin for password reset", zap.Error(err))
		return err
	}

	// Find user by email or username
	params := gocloak.GetUsersParams{
		Email:    &identifier,
		Username: &identifier,
	}
	users, err := s.kc.GetUsers(ctx, token.AccessToken, realm, params)
	if err != nil {
		s.log.Error("Failed to get user by identifier", zap.Error(err))
		return err
	}

	if len(users) == 0 {
		s.log.Error("User not found", zap.String("identifier", identifier))
		return fmt.Errorf("user not found")
	}

	userID := *users[0].ID

	// Reset password
	err = s.kc.SetPassword(ctx, token.AccessToken, userID, realm, newPassword, false)
	if err != nil {
		s.log.Error("Failed to reset password", zap.Error(err))
		return err
	}

	s.log.Info("Password reset successfully", zap.String("identifier", identifier))
	return nil
}

func (s *Service) GetUsernameByEmail(ctx context.Context, email string) (string, error) {
	// retrieve admin token
	realm := s.cfg.Keycloak.Realm
	token, err := s.kc.LoginAdmin(ctx, s.cfg.Keycloak.User, s.cfg.Keycloak.Passwd, realm)
	if err != nil {
		s.log.Error("Failed to login admin for getting username", zap.Error(err))
		return "", err
	}

	// Find user by email
	params := gocloak.GetUsersParams{
		Email: &email,
	}
	users, err := s.kc.GetUsers(ctx, token.AccessToken, realm, params)
	if err != nil {
		s.log.Error("Failed to get user by email", zap.Error(err))
		return "", err
	}

	if len(users) == 0 {
		s.log.Error("User not found", zap.String("email", email))
		return "", fmt.Errorf("user not found")
	}

	username := *users[0].Username
	s.log.Debug("Username retrieved", zap.String("email", email), zap.String("username", username))
	return username, nil
}

func getStringValue(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return "en"
}
