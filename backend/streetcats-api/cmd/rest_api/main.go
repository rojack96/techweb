package main

import (
	"context"
	"streetcats-api/api/server"
	"streetcats-api/api/server/router"
	"streetcats-api/configs"
	"streetcats-api/internal/repositories/vehicle"
	"streetcats-api/internal/services/email"
	"streetcats-api/pkg/eureka"
	"streetcats-api/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	@title			Notification Service
//	@version		0.1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@securityDefinitions.apiKey	Bearer
//	@in							header
//	@name						Authorization

//	@securityDefinitions.basic	BasicAuth
//	@in							header
//	@name						Basic Auth

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/notification-service/api/v1

// Order of tag
//	@tag.name	Notification

// @schemes	http
func main() {
	ctx := context.Background()
	sh := configs.NewServiceHub(ctx)

	// eureka configuration
	e := eureka.NewClient(sh.Config.Eureka.Enabled,
		eureka.WithMachines(sh.Config.Eureka.Machines),
		eureka.WithHostName(sh.Config.Eureka.HostName),
		eureka.WithHost(sh.Config.Eureka.Host, sh.Config.Eureka.Port),
		eureka.WithVipAddress(sh.Config.Eureka.VipAddress),
		eureka.WithSecureVipAddress(sh.Config.Eureka.SecureVipAddress),
		eureka.WithHomePageUrl(sh.Config.Eureka.HomePageUrl),
		eureka.WithHealthCheckUrl(sh.Config.Eureka.HealthCheckUrl),
		eureka.WithStatusPageUrl(sh.Config.Eureka.StatusPageUrl),
		eureka.WithDataCenterInfo(sh.Config.Eureka.DataCenterInfo),
		eureka.WithApp(sh.Config.Eureka.App),
		eureka.WithZapLogger(sh.Log),
	)

	fargo, err := e.BuildFargoInstance()
	if err != nil {
		sh.Log.Info("Failed to build fargo", zap.Error(err))
		panic(err)
	}
	fargo.Register()

	defer sh.Log.Sync() // ignore creation of error statement
	// gin config
	gin.SetMode(setGinMode(sh))
	zapWriter := &logger.ZapGinWriter{Logger: sh.Log}

	gin.DefaultWriter = zapWriter
	gin.DefaultErrorWriter = zapWriter

	// router configuration
	r := router.NewRouter(zapWriter, sh)

	// emailConfiguration
	vhRepo := vehicle.NewVehicleRepository(sh.PgCoreDbSession, sh.Log)
	emailService := email.NewEmailService(sh.Log, sh.RedisClient, sh.Email, sh.Config.Email, vhRepo)
	wk := email.NewEmailWorker(emailService, sh.RedisClient, sh.Log)

	go wk.Run(ctx)

	// server configuration
	sh.Log.Info("starting server configuration")
	s := server.NewServer(
		server.WithHost(sh.Config.Api.Host, sh.Config.Api.Port),
		server.WithRouter(r),
		server.WithSwagger(
			sh.Config.Api.Swagger.Enabled,
			"notification-service",
			sh.Config.Api.Swagger.Auth.Enabled,
			sh.Config.Api.Swagger.Auth.User, sh.Config.Api.Swagger.Auth.Passwd,
		),
		server.WithZapLogger(sh.Log),
	)

	s.Serve()
}

func setGinMode(sh *configs.ServiceHub) string {
	env := sh.GetEnvironment()
	if env == "prod" {
		return gin.ReleaseMode
	}
	return gin.DebugMode
}
