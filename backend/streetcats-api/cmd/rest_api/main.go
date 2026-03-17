package main

import (
	"context"
	"streetcats-api/api/server"
	"streetcats-api/api/server/router"
	"streetcats-api/configs"
	"streetcats-api/pkg/logger"

	"github.com/gin-gonic/gin"
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
	sh, err := configs.NewServiceHub(ctx)
	if err != nil {
		panic(err)
	}

	defer sh.Log.Sync() // ignore creation of error statement
	// gin config
	gin.SetMode("debug")
	zapWriter := &logger.ZapGinWriter{Logger: sh.Log}

	gin.DefaultWriter = zapWriter
	gin.DefaultErrorWriter = zapWriter

	// router configuration
	r := router.NewRouter(zapWriter, sh)

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
