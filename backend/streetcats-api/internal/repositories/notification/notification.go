package notification

import (
	"sipli/notification-service/internal/entities"
)

type Repository interface {
	CreateNotificationSchemas() error
	CreateNotificationsTables() error
	GetUserIdByPreferredUsername(username string) (uint64, error)
	GetRulesByImei(source, category string) ([]entities.RuleCustomer, error)
	GetRulesByVehicleId(source, category string) ([]entities.RuleCustomer, error)
	GetRulesByCustomerId(source, category string) ([]entities.RuleCustomer, error)
	GetUsersByRuleId(ruleId uint64) ([]entities.UserInfo, error)
	GetRecipientsByRuleId(ruleId uint64) ([]entities.UserInfo, error)
	GetUsersIdByImei(customerId uint64, source string) ([]entities.UserInfo, error)
	GetUsersIdByVehicleId(customerId uint64, source string) ([]entities.UserInfo, error)
	GetUsersIdByCustomerId(customerId uint64) ([]entities.UserInfo, error)
	GetRecipientsByCustomerId(customerId uint64) ([]entities.EmailReceipt, error)
	GetNotificationsByUserId(userId uint64) ([]entities.Notification, error)
	UpdateEventByUserId(userId uint64, events []uint64, markerState bool) error
	SaveAlert(alert entities.Alert) error
	SaveAlertTx(tx any, alert entities.Alert) error
	SaveEvent(event entities.Event) (uint64, error)
	SaveEventTx(tx any, event entities.Event) (uint64, error)
	SaveEventUser(eventUsers []entities.EventUser) error
	SaveEventUserTx(tx any, eventUsers []entities.EventUser) error
}
