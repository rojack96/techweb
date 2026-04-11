package router

import (
	// go imports
	// project imports
	"streetcats-api/configs"
	"streetcats-api/internal/controllers"
	ac "streetcats-api/internal/controllers/auth"
	sc "streetcats-api/internal/controllers/sightings"
	sr "streetcats-api/internal/repositories/sightings"
	"streetcats-api/internal/repositories/users"
	kcs "streetcats-api/internal/services/keycloak"
	sess "streetcats-api/internal/services/session"
	ss "streetcats-api/internal/services/sightings"
	us "streetcats-api/internal/services/users"

	// external imports
	"github.com/gin-gonic/gin"
)

type Register struct {
	publicRouter    *gin.RouterGroup
	protectedRouter *gin.RouterGroup
	sh              *configs.ServiceHub
}

func NewRegister(publicRouter *gin.RouterGroup, protectedRouter *gin.RouterGroup, sh *configs.ServiceHub) *Register {
	return &Register{publicRouter: publicRouter, protectedRouter: protectedRouter, sh: sh}
}

func (r *Register) AuthRoutes() {

	kcService := kcs.NewService(r.sh.Log, r.sh.Config, r.sh.Keycloak)
	sessionService := sess.NewService(r.sh.Log, r.sh.Config, r.sh.Keycloak, r.sh.RedisClient)
	controller := ac.NewAuthController(r.sh.Log, kcService, sessionService)

	authPublicGroup := r.publicRouter.Group("/auth")

	authPublicGroup.POST("/login", controller.Login)

	authProtectedGroup := r.protectedRouter.Group("/auth")
	authProtectedGroup.GET("/me", controller.Me)
	authProtectedGroup.POST("/logout", controller.Logout)

}

func (r *Register) UserRoutes() {

	usersRepo := users.NewUsersRepository(r.sh.Postgis)
	usersService := us.NewUsersService(r.sh.Log, r.sh.Config, r.sh.Keycloak, usersRepo)
	controller := controllers.NewController(r.sh.Log, usersService)

	userGroup := r.publicRouter.Group("/user")

	userGroup.POST("/register", controller.RegisterUser)
	userGroup.POST("/reset-password", controller.ResetPassword)
	userGroup.GET("/username/:email", controller.GetUsernameByEmail)
}

func (r *Register) SightingRoutes() {
	sightingsRepo := sr.NewRepository(r.sh.Postgis)
	sightingsService := ss.NewService(r.sh.Log, r.sh.Config, r.sh.Keycloak, sightingsRepo)
	controller := sc.NewController(r.sh.Log, sightingsService)

	sightingsGroup := r.publicRouter.Group("/sightings")
	sightingsGroup.GET("/:animalID/all", controller.AllSightings)
}
