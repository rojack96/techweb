package router

import (
	// go imports
	// project imports
	"streetcats-api/configs"
	"streetcats-api/internal/controllers"
	sc "streetcats-api/internal/controllers/sightings"
	sr "streetcats-api/internal/repositories/sightings"
	"streetcats-api/internal/repositories/users"
	ss "streetcats-api/internal/services/sightings"
	us "streetcats-api/internal/services/users"

	// external imports
	"github.com/gin-gonic/gin"
)

type Register struct {
	router *gin.RouterGroup
	sh     *configs.ServiceHub
}

func NewRegister(router *gin.RouterGroup, sh *configs.ServiceHub) *Register {
	return &Register{router: router, sh: sh}
}

func (r *Register) UserRoutes() {

	usersRepo := users.NewUsersRepository(r.sh.Postgis)
	usersService := us.NewUsersService(r.sh.Log, r.sh.Config, r.sh.Keycloak, usersRepo)
	controller := controllers.NewController(r.sh.Log, usersService)

	userGroup := r.router.Group("/user")

	userGroup.POST("/register", controller.RegisterUser)
	userGroup.POST("/reset-password", controller.ResetPassword)
	userGroup.GET("/username/:email", controller.GetUsernameByEmail)
}

func (r *Register) SightingRoutes() {
	sightingsRepo := sr.NewRepository(r.sh.Postgis)
	sightingsService := ss.NewService(r.sh.Log, r.sh.Config, r.sh.Keycloak, sightingsRepo)
	controller := sc.NewController(r.sh.Log, sightingsService)

	sightingsGroup := r.router.Group("/sightings")
	sightingsGroup.GET("/:animalID/all", controller.AllSightings)
}
