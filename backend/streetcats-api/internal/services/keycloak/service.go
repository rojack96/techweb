package keycloak

import (
	"context"
	"fmt"
	"net/url"
	"streetcats-api/configs"
	"streetcats-api/internal/repositories/users"

	"github.com/Nerzal/gocloak/v13"
	"go.uber.org/zap"
)

type ServiceInterfaces interface {
	GetLoginURL(state string) string
	ExchangeCode(code string) (*gocloak.JWT, error)
	GetUserInfo(accessToken string) (*gocloak.UserInfo, error)
	RefreshToken(refreshToken string) (*gocloak.JWT, error)
}

type service struct {
	log            *zap.Logger
	cfg            *configs.ConfigModel
	kc             *gocloak.GoCloak
	ctx            context.Context
	userRepository users.Repository
}

func NewService(log *zap.Logger, cfg configs.ConfigModel, kc *gocloak.GoCloak) ServiceInterfaces {
	return &service{log: log, cfg: &cfg, kc: kc, ctx: context.Background()}
}

func (k *service) GetLoginURL(state string) string {

	baseUrl := fmt.Sprintf("%s:%d", k.cfg.Keycloak.Host, k.cfg.Keycloak.Port)
	redirectURI := fmt.Sprintf("%s:%d/auth/callback", k.cfg.Api.Host, k.cfg.Api.Port)

	authURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/auth",
		baseUrl,
		k.cfg.Keycloak.Realm,
	)

	params := url.Values{}
	params.Add("client_id", k.cfg.Keycloak.ClientId)
	params.Add("response_type", "code")
	params.Add("redirect_uri", redirectURI)
	params.Add("scope", "openid profile email")
	params.Add("state", state)

	return fmt.Sprintf("%s?%s", authURL, params.Encode())
}

func (k *service) ExchangeCode(code string) (*gocloak.JWT, error) {
	redirectURI := fmt.Sprintf("%s:%d/auth/callback", k.cfg.Api.Host, k.cfg.Api.Port)
	grantType := "authorization_code"
	options := gocloak.TokenOptions{
		ClientID:     &k.cfg.Keycloak.ClientId,
		ClientSecret: &k.cfg.Keycloak.ClientSecret,
		GrantType:    &grantType,
		Code:         &code,
		RedirectURI:  &redirectURI,
	}
	return k.kc.GetToken(k.ctx, k.cfg.Keycloak.Realm, options)
}

func (k *service) GetUserInfo(accessToken string) (*gocloak.UserInfo, error) {
	return k.kc.GetUserInfo(k.ctx, accessToken, k.cfg.Keycloak.Realm)
}

func (k *service) RefreshToken(refreshToken string) (*gocloak.JWT, error) {
	return k.kc.RefreshToken(k.ctx, refreshToken, k.cfg.Keycloak.ClientId, k.cfg.Keycloak.ClientSecret, k.cfg.Keycloak.Realm)
}
