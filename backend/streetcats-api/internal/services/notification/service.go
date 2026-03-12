package notification

import (
	"context"
	channelhandler "sipli/notification-service/api/server/channel_handler"
	"sipli/notification-service/internal/dto"
	"sipli/notification-service/internal/infrastructure/persistance/gorm"
	"sipli/notification-service/internal/repositories/customer"
	"sipli/notification-service/internal/repositories/notification"
	"sipli/notification-service/internal/repositories/vehicle"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	NoSourceFounded = "no source founded"
)

type ServiceInterfaces interface {
	CreateNotificationSchemas() error
	CreateEventTables() error
	SaveEvent(r *channelhandler.ChannelHandler, request dto.NotificationEventDTO) ([]dto.EmailInfoDTO, error)
	UpdateEventByUserId(claims any, events []uint64, markerState bool) error
	GetNotifications(userId uint64) ([]dto.NotificationDTO, error)
	MappingMessage(message dto.NotificationEventDTO) (dto.NotificationDTO, error)
	GetUserIdByPreferredUsername(claims any) (uint64, error)
}

type Service struct {
	notificationRepo notification.Repository
	vehicleRepo      vehicle.Repository
	customerRepo     customer.Repository
	rds              *redis.Client
	log              *zap.Logger
	ctx              context.Context
	txManager        gorm.TransactionManager
}

func (s *Service) SetContext(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()
	s.ctx = reqCtx
}

func NewNotificationService(
	log *zap.Logger,
	notificationRepo notification.Repository,
	vehicleRepo vehicle.Repository,
	customerRepo customer.Repository,
	rds *redis.Client,
	txManager gorm.TransactionManager,
) *Service {
	return &Service{
		log:              log,
		notificationRepo: notificationRepo,
		vehicleRepo:      vehicleRepo,
		customerRepo:     customerRepo,
		rds:              rds,
		txManager:        txManager,
	}
}
