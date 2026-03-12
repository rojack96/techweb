package postgres

import (
	// go import
	"context"
	"fmt"

	// external import
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	// project import
)

const LogPrefix = "[POSTGRES]"

type Client struct {
	enabled  bool
	host     string
	user     string
	passwd   string
	dbName   string
	port     uint16
	log      *zap.Logger
	pgLogger struct {
		enabled  bool
		logLevel string
	}
}

func NewClient(enabled bool, opts ...Options) *Client {
	c := &Client{
		enabled: enabled,
		host:    "localhost",
		port:    5432,
		dbName:  "postgres",
		pgLogger: struct {
			enabled  bool
			logLevel string
		}{enabled: false, logLevel: "INFO"},
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Options - designed follow the idiomatic Go Functional options pattern
type Options func(*Client)

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

func WithDbName(dbName string) Options {
	return func(c *Client) {
		c.dbName = dbName
	}
}

func WithZapLogger(log *zap.Logger) Options {
	return func(c *Client) {
		c.log = log
	}
}

func WithPgLogLevel(enabled bool, level string) Options {
	return func(c *Client) {
		c.pgLogger.enabled = enabled
		c.pgLogger.logLevel = level
	}
}

// Connect establishes a pgx connection pool to PostgreSQL and returns it.
// It ignores logger options that were previously used by gorm.
func (pg *Client) Connect() (*pgxpool.Pool, error) {
	if !pg.enabled {
		pg.log.Warn(LogPrefix + " service is disabled!")
		return nil, nil
	}

	pg.log.Debug(LogPrefix + " service enabled")

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		pg.user, pg.passwd, pg.host, pg.port, pg.dbName,
	)

	// parse config to ensure it's valid (and to allow future tuning)
	_, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		pg.log.Warn(LogPrefix+" pool config parse error", zap.Error(err))
		return nil, err
	}

	// TODO: attach zap logger to cfg.ConnConfig.Logger if needed

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		pg.log.Warn(LogPrefix+" pool not opened", zap.Error(err))
		return nil, err
	}

	pg.log.Debug(LogPrefix + " pool successfully started")
	pg.log.Info("database pool started")

	return pool, nil
}

// gormConfig and setLogLevel are no longer used since the client uses pgx.
// They are left here for reference but can be removed in a future cleanup.

// NOTE: if we decide to support pgx logging, we would add methods to
// configure the pgxpool.Config.Logger field.
