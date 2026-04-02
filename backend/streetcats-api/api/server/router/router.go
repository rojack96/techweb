package router

import (
	// go imports
	"net/http"
	// external imports
	"github.com/gin-gonic/gin"
	// project imports
	"streetcats-api/api/server/middlewares"
	"streetcats-api/configs"
	"streetcats-api/pkg/logger"
)

func NewRouter(zapWriter *logger.ZapGinWriter, sh *configs.ServiceHub) *gin.Engine {
	r := gin.New() // creare a router without default middleware
	// todo check if is necessary
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		return nil
	}
	r.Use(
		gin.LoggerWithWriter(zapWriter),
		gin.RecoveryWithWriter(zapWriter),
	)

	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{"message": "Command API Middleware for Dev and PreProd"},
		)
	})

	//handlers
	const PrefixPath = "/streetcats-service/api/v1/"
	apiGroup := r.Group(PrefixPath)
	if sh.Config.Keycloak.Enabled {
		apiGroup.Use(middlewares.Auth(sh.Config, sh.Keycloak, PrefixPath))
	}

	register := NewRegister(apiGroup, sh)
	register.UserRoutes()
	register.SightingRoutes()

	return r
}
