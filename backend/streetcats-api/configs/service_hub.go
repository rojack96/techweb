package configs

import (
	// go import
	"context"

	// external import
	"github.com/Nerzal/gocloak/v13"
	//r "github.com/redis/go-redis/v9"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	// project import
	"streetcats-api/pkg/keycloak"
	"streetcats-api/pkg/logger"
	"streetcats-api/pkg/postgres"
)

type ServiceHub struct {
	environment string
	Config      ConfigModel
	Log         *zap.Logger
	Postgis     *pgxpool.Pool
	//	RedisClient *r.Client
	Keycloak *gocloak.GoCloak
}

func NewServiceHub(ctx context.Context) (*ServiceHub, error) {
	var (
		s   ServiceHub
		err error
	)

	s.Config, err = Config()
	if err != nil {
		return nil, err
	}

	// logger configuration
	s.Log = logger.NewZapLogger(s.Config.Log.Level, s.environment,
		logger.WithServiceName("streetcats-service")).
		Init()
	s.Log.Debug(logger.PrefixServerLog + " - ServiceHub init")
	s.Log.Info("starting all service configurations")
	// postgres configuration
	pg := postgres.NewClient(
		s.Config.Postgis.Enabled,
		postgres.WithHost(s.Config.Postgis.Host, s.Config.Postgis.Port),
		postgres.WithDbName(s.Config.Postgis.DbName),
		postgres.WithAuth(s.Config.Postgis.User, s.Config.Postgis.Passwd),
		postgres.WithZapLogger(s.Log),
		postgres.WithPgLogLevel(s.Config.Postgis.Logger.Enable, s.Config.Postgis.Logger.Level),
	)

	if s.Postgis, err = pg.Connect(); err != nil {
		s.Log.Info("Failed to connect to Postgres", zap.Error(err))
		panic(err)
	}

	// redis configuration
	/*rds := redis.NewClient(s.Config.RedisDb.Enabled,
		redis.WithHost(s.Config.RedisDb.Host, s.Config.RedisDb.Port),
		redis.WithDb(5),
		redis.WithAuth("default", s.Config.RedisDb.Passwd),
		redis.WithZapLogger(s.Log),
	)

	if s.RedisClient, err = rds.Connect(); err != nil {
		s.Log.Info("Failed to connect to Redis", zap.Error(err))
		panic(err)
	}*/

	// keycloak configuration
	kc := keycloak.NewClient(s.Config.Keycloak.Enabled,
		keycloak.WithHost(s.Config.Keycloak.Host, s.Config.Keycloak.Port),
		keycloak.WithClientId(s.Config.Keycloak.ClientId),
		keycloak.WithClientSecret(s.Config.Keycloak.ClientSecret),
		keycloak.WithRealm(s.Config.Keycloak.Realm),
		keycloak.WithZapLogger(s.Log),
		keycloak.WithCtx(ctx),
	)

	if s.Keycloak, err = kc.Connect(); err != nil {
		s.Log.Info("Failed to connect to Keycloak", zap.Error(err))
		panic(err)
	}

	return &s, nil
}
