package notification

import (
	"errors"
	"sipli/notification-service/internal/entities"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type notificationRepositoryImpl struct {
	pgCore *gorm.DB
	log    *zap.Logger
}

func (n *notificationRepositoryImpl) UpdateEventByUserId(userId uint64, events []uint64, markerState bool) error {
	if err := n.pgCore.
		Model(&entities.EventUser{}).
		Where("id_user = ?", userId).
		Where("id_event IN ?", events).
		Update("marked_state", markerState).Error; err != nil {
		return err
	}
	return nil
}

func (n *notificationRepositoryImpl) GetNotificationsByUserId(userId uint64) ([]entities.Notification, error) {
	var results []entities.Notification

	if err := n.pgCore.
		Select(`
			e.id,
			COALESCE(imei_vh.licenseplate, vehicle_source.licenseplate) AS license_plate,
			customer_source.companyname AS customer_name,
			user_source.name AS ur_name,
			user_source.surname AS ur_surname,
			user_source.username AS ur_username,
			e.data,
			e.ts,
			e.category,
			e.severity,
			e.alert_code,
			eu.marked_state`).
		Table("notification.event e").
		Joins("JOIN notification.event_user eu ON e.id = eu.id_event").
		// Take vehicle id from IMEI
		Joins("LEFT JOIN crmdb.ws_units wu ON e.source_type = 1 AND wu.imei = e.source").
		Joins(" LEFT JOIN crmdb.dv_devices dv ON dv.idunit = wu.id").
		Joins("LEFT JOIN crmdb.vh_vehicle imei_vh ON imei_vh.id = dv.idvehicle").
		// Take vehicle id from vehicle table
		Joins("LEFT JOIN crmdb.vh_vehicle vehicle_source ON e.source_type = 2 AND vehicle_source.id = e.source::bigint").
		// Take customer information
		Joins("LEFT JOIN crmdb.cl_clients customer_source ON e.source_type = 3 AND customer_source.id = e.source::bigint").
		// Take user information
		Joins("LEFT JOIN crmdb.ur_users user_source ON e.source_type = 4 AND user_source.id = e.source::bigint").
		Where("eu.id_user = ?", userId).
		Order("e.ts DESC").
		Limit(200).Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (n *notificationRepositoryImpl) GetRulesByCustomerId(source, category string) ([]entities.RuleCustomer, error) {
	var result []entities.RuleCustomer

	query := n.pgCore.Select("r.*, s.code AS severity_code, cl.id AS customer").
		Table("crmdb.cl_clients cl").
		Where("cl.id = ?", source)

	query = n.applyNotificationRule(query, category)

	if err := query.Scan(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (n *notificationRepositoryImpl) GetRulesByVehicleId(source, category string) ([]entities.RuleCustomer, error) {
	var result []entities.RuleCustomer

	query := n.pgCore.
		Select("r.*, s.code AS severity_code, cl.id AS customer").
		Table("crmdb.vh_vehicle vh").
		Joins("JOIN crmdb.cl_clients cl ON cl.id = vh.idclient").
		Where("vh.id = ?", source)

	query = n.applyNotificationRule(query, category)

	if err := query.Scan(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (n *notificationRepositoryImpl) GetRulesByImei(source, category string) ([]entities.RuleCustomer, error) {
	var result []entities.RuleCustomer

	query := n.pgCore.
		Select("r.*, s.code AS severity_code,cl.id AS customer").
		Table("crmdb.ws_units wu").
		Joins("JOIN crmdb.dv_devices dv ON dv.idunit = wu.id").
		Joins("JOIN crmdb.vh_vehicle vh ON vh.id = dv.idvehicle").
		Joins("JOIN crmdb.cl_clients cl ON cl.id = vh.idclient").
		Where("wu.imei = ?", source)

	query = n.applyNotificationRule(query, category)

	if err := query.Scan(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (n *notificationRepositoryImpl) GetRecipientsByRuleId(ruleId uint64) ([]entities.UserInfo, error) {
	var results []entities.UserInfo

	if err := n.pgCore.
		Select("erec.id, erec.name, erec.surname, erec.email, lang.code language").
		Table("notification.email_receipt erec").
		Joins("LEFT JOIN crmdb.ss_languages as lang on erec.language_id = lang.id").
		Joins("JOIN notification.email_rule er on er.recipient_id = erec.id").
		Where("er.rule_id = ?", ruleId).
		Where("erec.is_active").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (n *notificationRepositoryImpl) GetUsersByRuleId(ruleId uint64) ([]entities.UserInfo, error) {
	var results []entities.UserInfo

	err := n.pgCore.
		Select("ur.id usr_id, ur.name, ur.surname, ur.email, ur.username, ur.iduserrole, lang.code language").
		Table("crmdb.ur_users ur").
		Joins("LEFT JOIN crmdb.ss_languages as lang on ur.idlanguage = lang.id").
		Joins("JOIN notification.user_rule usrule ON usrule.user_id = ur.id").
		Joins("JOIN notification.rule r ON usrule.rule_id = r.id").
		Where("r.id = ?", ruleId).Where("ur.isactive").Where("r.is_active").
		Scan(&results).Error

	if err != nil {
		n.log.Debug("error occurred in query GetUsersByRuleId", zap.Error(err))
		return nil, err
	}

	return results, nil
}

func (n *notificationRepositoryImpl) SaveAlertTx(tx any, alert entities.Alert) error {
	gormTx, ok := tx.(*gorm.DB)
	if !ok {
		return errors.New("invalid transaction")
	}

	if err := gormTx.Create(&alert).Error; err != nil {
		return err
	}
	return nil
}

func (n *notificationRepositoryImpl) SaveAlert(alert entities.Alert) error {
	if err := n.pgCore.Create(&alert).Error; err != nil {
		return err
	}
	return nil
}

func (n *notificationRepositoryImpl) SaveEventUserTx(tx any, eventUsers []entities.EventUser) error {
	gormTx, ok := tx.(*gorm.DB)
	if !ok {
		return errors.New("invalid transaction")
	}

	if err := gormTx.Create(&eventUsers).Error; err != nil {
		return err
	}

	return nil
}

func (n *notificationRepositoryImpl) SaveEventUser(eventUsers []entities.EventUser) error {
	tx := n.pgCore.Begin()

	if err := tx.Create(&eventUsers).Error; err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (n *notificationRepositoryImpl) SaveEventTx(tx any, event entities.Event) (uint64, error) {
	gormTx, ok := tx.(*gorm.DB)
	if !ok {
		return 0, errors.New("invalid transaction")
	}

	if err := gormTx.Create(&event).Error; err != nil {
		return 0, err
	}

	return event.ID, nil
}

// SaveEvent - Save an event
func (n *notificationRepositoryImpl) SaveEvent(event entities.Event) (uint64, error) {
	if err := n.pgCore.Create(&event).Error; err != nil {
		return 0, err
	}

	return event.ID, nil
}

// CreateNotificationsTables - Create Event, EventUser tables
func (n *notificationRepositoryImpl) CreateNotificationsTables() error {
	if err := n.pgCore.AutoMigrate(&entities.Event{}, &entities.EventUser{}); err != nil {
		return err
	}
	return nil
}

// GetUserIdByPreferredUsername - Retrieve id using username provided
func (n *notificationRepositoryImpl) GetUserIdByPreferredUsername(username string) (uint64, error) {
	var user entities.UrUsers
	if err := n.pgCore.Select("id").
		Where("username", username).First(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (n *notificationRepositoryImpl) CreateNotificationSchemas() error {
	if err := n.pgCore.Exec("CREATE SCHEMA IF NOT EXISTS " + entities.NotificationSchema).Error; err != nil {
		return err
	}
	return nil
}

func (n *notificationRepositoryImpl) GetUsersIdByCustomerId(customerId uint64) ([]entities.UserInfo, error) {
	var results []entities.UserInfo
	if err := n.pgCore.
		Select("usr.id usr_id, usr.name, usr.surname, usr.email, usr.username, usr.iduserrole, lang.code as language").
		Table("crmdb.ur_users AS usr").
		Joins("LEFT JOIN crmdb.ss_languages as lang on usr.idlanguage = lang.id").
		Where("usr.idclient = ?", customerId).
		//Distinct().
		Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (n *notificationRepositoryImpl) GetRecipientsByCustomerId(customerId uint64) ([]entities.EmailReceipt, error) {
	var receipts []entities.EmailReceipt

	if err := n.pgCore.
		Select("er.id, er.name, er.surname, er.email, sl.code as language").
		Table("notification.email_receipt er").
		Joins("LEFT JOIN crmdb.ss_languages sl on er.language_id = sl.id").
		Where("er.customer_id = ?", customerId).
		Scan(&receipts).Error; err != nil {
		return nil, err
	}

	return receipts, nil
}

func (n *notificationRepositoryImpl) GetUsersIdByVehicleId(customerId uint64, source string) ([]entities.UserInfo, error) {
	var results []entities.UserInfo

	if err := n.pgCore.
		Select("usr.id usr_id, usr.name, usr.surname, usr.email, usr.username, usr.iduserrole, usvh.id usvhid, lang.code as language").
		Table("crmdb.ur_users usr").
		Joins("LEFT JOIN crmdb.ss_languages as lang on usr.idlanguage = lang.id").
		Joins("JOIN crmdb.cl_clients AS cl ON usr.idclient = cl.id").
		Joins("JOIN crmdb.vh_vehicle AS vh ON vh.idclient = cl.id").
		Joins("LEFT JOIN crmdb.ur_user_vehicles AS usvh ON usvh.iduser = usr.id AND usvh.idvehicle = vh.id").
		Where("usr.idclient = ?", customerId).
		Where("vh.id = ?", source).
		Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (n *notificationRepositoryImpl) GetUsersIdByImei(customerId uint64, source string) ([]entities.UserInfo, error) {
	var results []entities.UserInfo

	if err := n.pgCore.
		Select("usr.id usr_id, usr.name, usr.surname, usr.email, usr.username, usr.iduserrole, usvh.id usvhid, lang.code as language").
		Table("crmdb.ur_users usr").
		Joins("LEFT JOIN crmdb.ss_languages as lang on usr.idlanguage = lang.id").
		Joins("JOIN crmdb.cl_clients AS cl ON usr.idclient = cl.id").
		Joins("JOIN crmdb.vh_vehicle AS vh ON vh.idclient = cl.id").
		Joins("JOIN crmdb.dv_devices AS dv ON dv.idvehicle = vh.id").
		Joins("JOIN crmdb.ws_units AS wu ON wu.id = dv.idunit").
		Joins("LEFT JOIN crmdb.ur_user_vehicles AS usvh ON usvh.iduser = usr.id AND usvh.idvehicle = vh.id").
		Where("usr.idclient = ?", customerId).
		Where("wu.imei = ?", source).
		Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func NewNotificationRepository(pgCore *gorm.DB, log *zap.Logger) Repository {
	return &notificationRepositoryImpl{pgCore: pgCore, log: log}
}

/* --------------- Utility function ---------------*/

func (n *notificationRepositoryImpl) applyNotificationRule(query *gorm.DB, category string) *gorm.DB {
	return query.
		Joins(`
		LEFT JOIN notification.rule r
		   ON r.customer_id = cl.id
			   AND r.category = ? 
               AND r.is_active = true

		LEFT JOIN notification.rule_severity rs
			ON rs.rule_id = r.id

		LEFT JOIN notification.severity s
			ON s.code = rs.severity_code
               AND s.is_active = true
		`, category)
}
