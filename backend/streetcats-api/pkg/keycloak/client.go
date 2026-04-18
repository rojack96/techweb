package keycloak

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"go.uber.org/zap"
)

type Client struct {
	enabled      bool
	host         string
	port         uint16
	clientId     string
	clientSecret string
	realm        string
	ctx          context.Context
	log          *zap.Logger
}

const LogPrefix = "[KEYCLOAK]"

func NewClient(enabled bool, opts ...Options) *Client {
	c := &Client{
		enabled: enabled,
		host:    "localhost",
		port:    9443,
		realm:   "master",
		ctx:     context.Background(),
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

type Options func(*Client)

func WithHost(host string, port uint16) Options {
	return func(c *Client) {
		c.host = host
		c.port = port
	}
}

func WithClientId(clientId string) Options {
	return func(c *Client) {
		c.clientId = clientId
	}
}

func WithClientSecret(clientSecret string) Options {
	return func(c *Client) {
		c.clientSecret = clientSecret
	}
}

func WithRealm(realm string) Options {
	return func(c *Client) {
		c.realm = realm
	}
}

func WithZapLogger(logger *zap.Logger) Options {
	return func(c *Client) {
		c.log = logger
	}
}

func WithCtx(ctx context.Context) Options {
	return func(c *Client) {
		c.ctx = ctx
	}
}

// TODO change all logs
func (kc *Client) Connect() (*gocloak.GoCloak, error) {
	var (
		err error
	)

	if !kc.enabled {
		kc.log.Warn(LogPrefix + " service is disabled!")
		return nil, nil
	}

	kc.log.Debug(LogPrefix + " service enabled")

	host := fmt.Sprintf("%s:%d/auth", kc.host, kc.port)
	client := gocloak.NewClient(host)

	if _, err = client.LoginClient(kc.ctx, kc.clientId, kc.clientSecret, kc.realm); err != nil {
		kc.log.Warn(LogPrefix + "connection not established!")
		return nil, err
	}
	kc.log.Debug(LogPrefix + " connection established")
	kc.log.Info(LogPrefix + " service connection established")

	return client, nil
}
