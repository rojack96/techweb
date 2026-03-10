package router

import (
	// go imports
	// project imports
	channelhandler "sipli/notification-service/api/server/channel_handler"
	"sipli/notification-service/configs"
	"sipli/notification-service/internal/controllers"
	"sipli/notification-service/internal/infrastructure/persistance/gorm"
	"sipli/notification-service/internal/repositories/alert"
	"sipli/notification-service/internal/repositories/customer"
	"sipli/notification-service/internal/repositories/notification"
	"sipli/notification-service/internal/repositories/vehicle"
	"sipli/notification-service/internal/services"
	"sipli/notification-service/internal/services/email"
	ntfService "sipli/notification-service/internal/services/notification"
	"time"

	// external imports
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Register struct {
	router *gin.RouterGroup
	sh     *configs.ServiceHub
}

func NewRegister(router *gin.RouterGroup, sh *configs.ServiceHub) *Register {
	return &Register{
		router: router,
		sh:     sh,
	}
}

func (r *Register) AlertRoutes() {
	repo := alert.NewAlertRepository(r.sh.PgCoreDbSession)
	service := services.NewAlertService(repo)
	// controllers := ....
	if err := service.CreateAlertSchemas(); err != nil {
		r.sh.Log.Error("Failed to create alert schemas", zap.Error(err))
		return
	}
}

func (r *Register) NotificationRoutes() {
	const maxConns = 10
	sendTimeout := 1 * time.Second
	ch := channelhandler.NewChannelHandler(r.sh.Log, maxConns, sendTimeout)

	ntfRepo := notification.NewNotificationRepository(r.sh.PgCoreDbSession, r.sh.Log)
	vhRepo := vehicle.NewVehicleRepository(r.sh.PgCoreDbSession, r.sh.Log)
	clRepo := customer.NewCustomerRepository(r.sh.PgCoreDbSession, r.sh.Log)
	emailService := email.NewEmailService(r.sh.Log, r.sh.RedisClient, r.sh.Email, r.sh.Config.Email, vhRepo)
	txManager := gorm.NewGormTransactionManager(r.sh.PgCoreDbSession)
	notificationService := ntfService.NewNotificationService(r.sh.Log, ntfRepo, vhRepo, clRepo, r.sh.RedisClient, txManager)
	controller := controllers.NewController(notificationService, emailService, r.sh.Log)

	if err := notificationService.CreateNotificationSchemas(); err != nil {
		r.sh.Log.Error("Failed to create notification schema", zap.Error(err))
		return
	}
	if err := notificationService.CreateEventTables(); err != nil {
		r.sh.Log.Error("Failed to create notification tables", zap.Error(err))
		return
	}

	notificationGroup := r.router.Group("/notification")
	notificationGroup.POST("/event", func(ctx *gin.Context) {
		controller.NotificationSentEvent(ctx, ch)
	})

	notificationGroup.GET("/stream", func(c *gin.Context) {
		controller.NotificationStream(c, ch)
	})

	notificationGroup.PATCH("/mark-event", controller.NotificationMarkEvent)
}
