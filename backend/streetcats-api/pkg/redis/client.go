package redis

import (
	"errors"
	"fmt"

	r "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Client struct {
	enabled bool
	host    string
	port    uint16
	user    string
	passwd  string
	db      int
	log     *zap.Logger
}

// Options - designed follow the idiomatic Go Functional options pattern
type Options func(*Client)

func NewClient(enabled bool, opts ...Options) *Client {
	c := &Client{
		enabled: enabled,
		host:    "localhost",
		port:    6379,
		db:      0,
		user:    "",
		passwd:  "",
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithHost(host string, port uint16) Options {
	return func(c *Client) {
		c.host = host
		c.port = port
	}
}

func WithAuth(user, passwd string) Options {
	return func(c *Client) {
		c.user = user
		c.passwd = passwd
	}
}

func WithDb(db int) Options {
	return func(c *Client) {
		c.db = db
	}
}

func WithZapLogger(log *zap.Logger) Options {
	return func(c *Client) {
		c.log = log
	}
}

func (c *Client) Connect() (*r.Client, error) {
	if !c.enabled {
		c.log.Warn("redis disabled")
		return nil, nil
	}

	addr := fmt.Sprintf("%s:%d", c.host, c.port)

	rds := r.NewClient(&r.Options{
		Addr:     addr,
		Password: c.passwd,
		DB:       c.db,
	})

	if rds == nil {
		c.log.Error("connect to redis failed")
		return rds, errors.New("connect to redis failed")
	}

	c.log.Info("connect to redis success")
	return rds, nil
}
