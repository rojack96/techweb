package router

import (
	// go imports
	"net/http"
	"time"

	// external imports
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// project imports
	"streetcats-api/api/server/middlewares"
	"streetcats-api/configs"
	"streetcats-api/internal/services/session"
	"streetcats-api/pkg/logger"
)

func NewRouter(zapWriter *logger.ZapGinWriter, sh *configs.ServiceHub) *gin.Engine {
	r := gin.New() // creare a router without default middleware
	// todo check if is necessary
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		return nil
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		}, // frontend
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

	publicGroup := r.Group(PrefixPath)

	protectedGroup := r.Group(PrefixPath)
	if sh.Config.Keycloak.Enabled {
		sessionService := session.NewService(sh.Log, sh.Config, sh.Keycloak, sh.RedisClient)
		protectedGroup.Use(middlewares.Auth(sh.Config, sh.Keycloak, sessionService))
	}

	register := NewRegister(publicGroup, protectedGroup, sh)
	register.AuthRoutes()
	register.UserRoutes()
	register.SightingRoutes()

	return r
}
