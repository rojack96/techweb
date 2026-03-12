package notification

import (
	"encoding/json"
	"errors"
	"fmt"
	channelhandler "sipli/notification-service/api/server/channel_handler"
	"sipli/notification-service/internal/dto"
	"sipli/notification-service/internal/entities"
	"sipli/notification-service/internal/enums"
	"sipli/notification-service/pkg/logger"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

func (s *Service) UpdateEventByUserId(claims any, events []uint64, markerState bool) error {
	var (
		userId uint64
		err    error
	)

	if userId, err = s.GetUserIdByPreferredUsername(claims); err != nil {
		return err
	}

	return s.notificationRepo.UpdateEventByUserId(userId, events, markerState)
}

// SaveEvent - Save a incoming event
func (s *Service) SaveEvent(r *channelhandler.ChannelHandler, request dto.NotificationEventDTO) ([]dto.EmailInfoDTO, error) {

	var (
		eventId              uint64
		jsonData             []byte
		users, recipients    []dto.UserInfoDTO
		rules                []dto.Rule
		rulesState           *bool
		alertCapability      bool
		emailInfo            []dto.EmailInfoDTO
		ErrNoRulesBySeverity = errors.New("no rules by severity")
	)

	if data, err := json.Marshal(request.Data); err != nil {
		return nil, err
	} else {
		jsonData = data
	}

	createDate := time.Now().Unix()

	event := entities.Event{
		Source:     request.Source,
		SourceType: int32(request.SourceType),
		Category:   request.Category,
		Severity:   request.Severity,
		AlertCode:  request.AlertCode,
		Data:       jsonData,
		Ts:         request.Ts,
		CreateDate: createDate,
	}

	err := s.txManager.Transaction(func(tx any) error {

		var err error

		if eventId, err = s.notificationRepo.SaveEventTx(tx, event); err != nil {
			return err
		}

		if rulesState, rules, err = s.checkIfExistRulesBySource(request.Source, request.SourceType, request.Category, request.Severity); err != nil {
			return err
		}

		if rulesState == nil {
			s.log.Info("no notification to sent",
				zap.String("source", logger.MaskStringSource(request.Source)),
				zap.String("category", request.Category),
				zap.String("severity", request.Severity))
			return ErrNoRulesBySeverity
		}

		if users, recipients, alertCapability, err = s.getUsers(*rulesState, rules, request.Source, request.SourceType); err != nil {
			return err
		}

		if request.SourceType == enums.Imei || request.SourceType == enums.VehicleId {
			if err = s.checkAlertCapability(tx, alertCapability, event.Data, request.Source, request.SourceType, event.Ts, event.CreateDate, request.AlertCode, event.Category, event.Severity); err != nil {
				return err
			}
		}

		eventUsers := make([]entities.EventUser, 0, len(users))
		for _, u := range users {
			eventUsers = append(eventUsers, entities.EventUser{
				IdEvent:     eventId,
				IdUser:      u.ID,
				MarkedState: false,
			})
		}

		return s.notificationRepo.SaveEventUserTx(tx, eventUsers)
	})

	if err != nil {
		if errors.Is(err, ErrNoRulesBySeverity) {
			return nil, nil
		}
		return nil, err
	}

	// 🔔 DOPO COMMIT
	request.ID = eventId

	msgToSent, err := s.MappingMessage(request)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		r.SendToUser(u.ID, msgToSent)
		if u.Email != nil && *u.Email != "" {
			emailInfo = append(emailInfo, dto.EmailInfoDTO{
				Name:            u.Name,
				Surname:         u.Surname,
				Email:           *u.Email,
				Username:        u.Username,
				Language:        u.Language,
				NotificationDTO: msgToSent,
			})
		}
	}

	for _, rc := range recipients {
		emailInfo = append(emailInfo, dto.EmailInfoDTO{
			Name:            rc.Name,
			Surname:         rc.Surname,
			Email:           *rc.Email,
			Username:        rc.Username,
			Language:        rc.Language,
			NotificationDTO: msgToSent,
		})
	}

	return emailInfo, nil
}

/* --------------- Redis ---------------*/

func (s *Service) setRecipientsBySource(customerId uint64, recipients []dto.UserInfoDTO) error {
	var (
		err error
	)

	table := fmt.Sprintf("notification:recipients:%d", customerId)

	insert := s.rds.JSONSet(s.ctx, table, "$", recipients)

	s.log.Debug("notification setting", zap.Any("table", table))

	if err = insert.Err(); err != nil {
		s.log.Error("failed to save recipients", zap.Error(err))
		return err
	}

	ttl := time.Hour * 24 * 365 * 5
	ttlSet := s.rds.Expire(s.ctx, table, ttl)

	if ttlSet.Err() != nil {
		s.log.Error("failed to save recipients", zap.Error(ttlSet.Err()))
		return ttlSet.Err()
	}

	return nil
}

// getRecipientsBySourceFromCache - Retrieve information about specific customer if you have
func (s *Service) getRecipientsBySourceFromCache(customerId uint64) ([]dto.UserInfoDTO, error) {
	var (
		res         string
		redisResult [][]dto.UserInfoDTO
		result      []dto.UserInfoDTO
		err         error
	)

	table := fmt.Sprintf("notification:recipients:%d", customerId)

	s.log.Debug("notification getting recipients", zap.Any("table", table))

	if res, err = s.rds.JSONGet(s.ctx, table, "$").Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			// chiave non presente
			return nil, nil
		}
		s.log.Error("failed to get customer have_rules", zap.Error(err))
		return nil, err
	}

	// Redis JSON con "$" ritorna sempre un array

	if err = json.Unmarshal([]byte(res), &redisResult); err != nil {
		s.log.Error("failed to unmarshal customer have_rules", zap.Error(err))
		return nil, err
	}

	if len(redisResult) == 0 {
		return nil, nil
	}

	result = redisResult[0]

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func (s *Service) setUsersBySource(customerId uint64, source string, sourceType uint16, users []dto.UserInfoDTO) error {
	var (
		err error
	)

	table := fmt.Sprintf("notification:users:%d:%d:%s", customerId, sourceType, source)

	insert := s.rds.JSONSet(s.ctx, table, "$", users)

	s.log.Debug("notification setting", zap.Any("table", table))

	if err = insert.Err(); err != nil {
		s.log.Error("failed to save users", zap.Error(err))
		return err
	}

	ttl := time.Hour * 24 * 365 * 5
	ttlSet := s.rds.Expire(s.ctx, table, ttl)

	if ttlSet.Err() != nil {
		s.log.Error("failed to save users", zap.Error(ttlSet.Err()))
		return ttlSet.Err()
	}

	return nil
}

// getUsersBySource - Retrieve information about specific customer if you have
func (s *Service) getUsersBySourceFromCache(customerId uint64, source string, sourceType uint16) ([]dto.UserInfoDTO, error) {
	var (
		res         string
		redisResult [][]dto.UserInfoDTO
		result      []dto.UserInfoDTO
		err         error
	)

	table := fmt.Sprintf("notification:users:%d:%d:%s", customerId, sourceType, source)

	s.log.Debug("notification getting users", zap.Any("table", table))

	if res, err = s.rds.JSONGet(s.ctx, table, "$").Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			// chiave non presente
			return nil, nil
		}
		s.log.Error("failed to get customer have_rules", zap.Error(err))
		return nil, err
	}

	// Redis JSON con "$" ritorna sempre un array

	if err = json.Unmarshal([]byte(res), &redisResult); err != nil {
		s.log.Error("failed to unmarshal customer have_rules", zap.Error(err))
		return nil, err
	}

	if len(redisResult) == 0 {
		return nil, nil
	}

	result = redisResult[0]

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

/* --------------- Utility service ---------------*/

func (s *Service) getVehicleIdBySource(source string, sourceType uint16) (uint64, error) {
	var (
		vehicleId uint64
		vhInfo    entities.VhVehicle
		rdsResult string
		err       error
	)

	switch sourceType {
	case enums.Imei:
		// get from redis
		const namespace = "imei_to_vehicleid:"
		key := namespace + source

		rdsResult, err = s.rds.Get(s.ctx, key).Result()
		if err == nil {
			if vehicleId, err = strconv.ParseUint(rdsResult, 10, 64); err == nil {
				return vehicleId, nil
			}
		} else if !errors.Is(err, redis.Nil) {
			s.log.Error("redis unavailable, fallback to DB", zap.String("key", logger.MaskStringSource(key)), zap.Error(err))
		}
		// get from postgres
		if vhInfo, err = s.vehicleRepo.GetVehicleInfoByImei(source); err != nil {
			return 0, err
		}

		vehicleId = vhInfo.ID
		// save on redis
		ttl := time.Hour * 24 * 30 // TLL thirty days
		if err = s.rds.Set(s.ctx, key, vehicleId, ttl).Err(); err != nil {
			s.log.Error("failed to save vehicle_id", zap.String("key", logger.MaskStringSource(key)), zap.Error(err))
		}
	case enums.VehicleId:
		if vehicleId, err = strconv.ParseUint(source, 10, 64); err != nil {
			return 0, err
		}
	default:
		return 0, errors.New(NoSourceFounded)
	}

	return vehicleId, nil
}

// checkIfExistRulesBySource - return customer id and rules (if exists) based on source, category and severity
func (s *Service) checkIfExistRulesBySource(source string, sourceType uint16, category, severity string) (*bool, []dto.Rule, error) {
	// TODO add redis
	var (
		function    func(string, string) ([]entities.RuleCustomer, error)
		queryResult []entities.RuleCustomer
		result      []dto.Rule
		ok          bool
		err         error
	)

	// TODO ricordarsi che il problema è le rules con una sola riga
	// Select a query based on the source
	funcMap := map[uint16]func(string, string) ([]entities.RuleCustomer, error){
		enums.Imei:       s.notificationRepo.GetRulesByImei,
		enums.VehicleId:  s.notificationRepo.GetRulesByVehicleId,
		enums.CustomerId: s.notificationRepo.GetRulesByCustomerId,
		//enums.DistributorId: nil,
		//enums.GeneralAppId:  nil,
	}

	if function, ok = funcMap[sourceType]; !ok {
		return nil, nil, errors.New(NoSourceFounded)
	}

	if queryResult, err = function(source, category); err != nil {
		return nil, nil, err
	}

	for _, rule := range queryResult {
		result = append(result, dto.Rule{
			Customer:               rule.Customer,
			ID:                     rule.ID,
			SeverityCode:           rule.SeverityCode,
			AlertCapability:        rule.AlertCapability,
			AppCapability:          rule.AppCapability,
			EmailCapability:        rule.EmailCapability,
			CustomerId:             rule.CustomerId,
			EmailToRecipients:      rule.EmailToRecipients,
			FilterUserEnabled:      rule.FilterUserEnabled,
			FilterRecipientEnabled: rule.FilterRecipientEnabled,
			CreatedAt:              rule.CreatedAt,
			UpdatedAt:              rule.UpdatedAt,
			IsActive:               rule.IsActive,
			Category:               rule.Category,
		})
	}

	rulesState := true
	if queryResult[0].ID == 0 {
		rulesState = false
		return &rulesState, result, nil
	}

	res := result[:0] // copy to new slice
	for _, r := range result {
		// filter by severity
		if r.SeverityCode == severity {
			res = append(res, r)
		}
	}

	result = res

	if len(result) == 0 { // category exist by severity not
		return nil, result, nil
	}

	return &rulesState, result, nil
}

// getUsers - Retrieve users and recipients using customer id
func (s *Service) getUsers(rulesState bool, rules []dto.Rule, source string, sourceType uint16) (users []dto.UserInfoDTO, recipients []dto.UserInfoDTO, alertCapability bool, err error) {
	if !rulesState {
		return s.getUsersWithoutFilter(rules[0].Customer, source, sourceType)
	}
	// if the customer don't have rules, sent to all users under specific customer
	return s.getUsersIdBySourceAndRule(rules, source, sourceType)

}

// getUsersWithoutFilter -
func (s *Service) getUsersWithoutFilter(customerId uint64, source string, sourceType uint16) (users []dto.UserInfoDTO, recipients []dto.UserInfoDTO, alertCapability bool, err error) {
	if users, err = s.getUsersBySource(customerId, source, sourceType); err != nil {
		return nil, nil, false, err
	}

	if recipients, err = s.getRecipientsByCustomerId(customerId); err != nil {
		return nil, nil, false, err
	}

	return users, recipients, false, nil
}

// getUsersIdBySourceAndRule -
func (s *Service) getUsersIdBySourceAndRule(rules []dto.Rule, source string, sourceType uint16) (users []dto.UserInfoDTO, recipients []dto.UserInfoDTO, alertCapability bool, err error) {
	var (
		allUsers, allRecipients []dto.UserInfoDTO
	)

	for _, rule := range rules {
		if rule.AlertCapability {
			alertCapability = true
		}
		// when the rule is active but the filter user and recipients are disabled
		// take users without filters
		if !rule.FilterUserEnabled && !rule.FilterRecipientEnabled {
			if users, recipients, _, err = s.getUsersWithoutFilter(rule.Customer, source, sourceType); err != nil {
				return nil, nil, false, err
			}

			if !rule.AppCapability {
				allUsers = uniqueUsersByID(allUsers, users)
			}

			if !rule.EmailCapability {
				allRecipients = uniqueUsersByID(allRecipients, recipients)
			}
			continue
		}

		var userResult []dto.UserInfoDTO
		if userResult, err = s.filterUserEnabled(rule, source, sourceType); err != nil {
			return nil, nil, false, err
		}

		if !rule.EmailCapability {
			for i := range userResult {
				userResult[i].Email = nil
			}
		}

		if rule.AppCapability {
			allUsers = uniqueUsersByID(allUsers, userResult)
		}

		var recipientsResult []dto.UserInfoDTO
		if recipientsResult, err = s.filterRecipientEnabled(rule); err != nil {
			return nil, nil, false, err
		}
		if rule.EmailCapability {
			allRecipients = uniqueUsersByID(allRecipients, recipientsResult)
		}

		// remove the email from accounts if email is sent only to recipients
		if rule.EmailToRecipients {
			for i := range allUsers {
				allUsers[i].Email = nil
			}
		}
	}

	users = allUsers
	recipients = allRecipients

	return users, recipients, alertCapability, nil
}

// getUsersBySource - Retrieve users based on source and source type
func (s *Service) getUsersBySource(customerId uint64, source string, sourceType uint16) ([]dto.UserInfoDTO, error) {
	var (
		users   []entities.UserInfo
		results []dto.UserInfoDTO
		err     error
	)

	// get users from cache
	results, err = s.getUsersBySourceFromCache(customerId, source, sourceType)

	if len(results) > 0 && err == nil {
		return results, nil
	}

	switch sourceType {
	case enums.Imei:
		if users, err = s.notificationRepo.GetUsersIdByImei(customerId, source); err != nil {
			return nil, err
		}
	case enums.VehicleId:
		if users, err = s.notificationRepo.GetUsersIdByVehicleId(customerId, source); err != nil {
			return nil, err
		}
	case enums.CustomerId:
		if users, err = s.notificationRepo.GetUsersIdByCustomerId(customerId); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(NoSourceFounded)
	}

	for _, u := range users {
		results = append(results, dto.UserInfoDTO{
			ID:   u.ID,
			Name: u.Name, Surname: u.Surname,
			Email: &u.Email, Username: &u.Username,
			IdUserRole: &u.IdUserRole,
			Language:   u.Language,
		})
	}

	if err = s.setUsersBySource(customerId, source, sourceType, results); err != nil {
		s.log.Error("Failed to set users by source", zap.Error(err))
	}

	return results, nil
}

// getRecipientsByCustomerId - Retrieve recipients using customer id
func (s *Service) getRecipientsByCustomerId(customerId uint64) ([]dto.UserInfoDTO, error) {
	var (
		rec        []entities.EmailReceipt
		recipients []dto.UserInfoDTO
		err        error
	)

	recipients, err = s.getRecipientsBySourceFromCache(customerId)

	if len(recipients) > 0 && err == nil {
		return recipients, nil
	}

	if rec, err = s.notificationRepo.GetRecipientsByCustomerId(customerId); err != nil {
		return nil, err
	}

	for _, r := range rec {
		recipients = append(recipients, dto.UserInfoDTO{
			ID:       r.ID,
			Name:     r.Name,
			Surname:  r.Surname,
			Email:    &r.Email,
			Language: &r.Language,
		})
	}

	if err = s.setRecipientsBySource(customerId, recipients); err != nil {
		s.log.Error("Failed to set recipients by source", zap.Error(err))
	}

	return recipients, nil
}

func (s *Service) filterRecipientEnabled(rule dto.Rule) ([]dto.UserInfoDTO, error) {
	var (
		recipients []entities.UserInfo
		result     []dto.UserInfoDTO
		err        error
	)

	if !rule.FilterRecipientEnabled {
		if result, err = s.getRecipientsByCustomerId(rule.Customer); err != nil {
			return nil, err
		}
	} else {

		if recipients, err = s.notificationRepo.GetRecipientsByRuleId(rule.ID); err != nil {
			return nil, err
		}

		for _, rc := range recipients {
			result = append(result, dto.UserInfoDTO{ID: rc.ID, Name: rc.Name, Surname: rc.Surname, Email: &rc.Email, Language: rc.Language})
		}
	}

	return result, nil
}

func (s *Service) filterUserEnabled(rule dto.Rule, source string, sourceType uint16) ([]dto.UserInfoDTO, error) {
	var (
		result []dto.UserInfoDTO
		users  []entities.UserInfo
		err    error
	)
	if !rule.FilterUserEnabled {
		if result, err = s.getUsersBySource(rule.Customer, source, sourceType); err != nil {
			return nil, err
		}
	} else {
		// TODO add redis
		if users, err = s.notificationRepo.GetUsersByRuleId(rule.ID); err != nil {
			return nil, err
		}

		for _, ur := range users {
			result = append(result, dto.UserInfoDTO{
				ID: ur.ID, Name: ur.Name, Surname: ur.Surname,
				Email: &ur.Email, Username: &ur.Username, Language: ur.Language,
				IdUserRole: &ur.IdUserRole, UsVhId: ur.UsVhId,
			})
		}

	}

	return result, nil
}

func (s *Service) checkAlertCapability(tx any, capability bool, data datatypes.JSON, source string, sourceType uint16, ts, createDate int64, alertCode string, category, severity string) error {
	const by = "sysadmin"
	var (
		vehicleId  uint64
		valueAlert *string
		val        any
		ok         bool
		err        error
	)

	if !capability {
		return nil
	}

	if val, ok, err = getJSONKey(data, "value"); err == nil && ok {
		v := val.(string)
		valueAlert = &v
	}

	if vehicleId, err = s.getVehicleIdBySource(source, sourceType); err != nil {
		return err
	}

	alert := entities.Alert{
		VehicleId: vehicleId,
		Timestamp: ts,
		AlertCode: alertCode, Category: category, Severity: severity,
		Properties: data, Value: valueAlert,
		Completed: false, CreatedDate: createDate, CreatedBy: by, UpdatedDate: createDate, UpdatedBy: by,
	}

	if err = s.notificationRepo.SaveAlertTx(tx, alert); err != nil {
		return err
	}

	return nil
}

/* --------------- Utility function ---------------*/

// getJSONKey - extract key from json
func getJSONKey(j datatypes.JSON, key string) (any, bool, error) {
	var m map[string]any

	if err := json.Unmarshal(j, &m); err != nil {
		return nil, false, err
	}

	val, exists := m[key]
	return val, exists, nil
}

// uniqueUsersByID - delete duplicated users
func uniqueUsersByID(arrays ...[]dto.UserInfoDTO) []dto.UserInfoDTO {
	seen := make(map[uint64]dto.UserInfoDTO)

	for _, arr := range arrays {
		for _, u := range arr {
			if _, exists := seen[u.ID]; !exists {
				seen[u.ID] = u
			}
		}
	}

	result := make([]dto.UserInfoDTO, 0, len(seen))
	for _, u := range seen {
		result = append(result, u)
	}

	return result
}
