package router

import (
	// go imports
	// project imports
	"streetcats-api/configs"
	"streetcats-api/internal/controllers"
	"streetcats-api/internal/repositories/users"
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
}
