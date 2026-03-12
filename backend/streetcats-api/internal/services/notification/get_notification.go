package notification

import (
	"encoding/json"
	"fmt"
	"sipli/notification-service/internal/dto"
	"sipli/notification-service/internal/entities"
	"sipli/notification-service/internal/enums"
	"strconv"
	"strings"

	"gorm.io/datatypes"
)

func (s *Service) GetNotifications(userId uint64) ([]dto.NotificationDTO, error) {
	var (
		queryResult []entities.Notification
		result      []dto.NotificationDTO
		err         error
	)
	if queryResult, err = s.notificationRepo.GetNotificationsByUserId(userId); err != nil {
		return nil, err
	}

	for _, notification := range queryResult {
		ntf := dto.NotificationDTO{
			ID:          notification.ID,
			Category:    notification.Category,
			Severity:    notification.Severity,
			AlertCode:   notification.AlertCode,
			Data:        notification.Data,
			Ts:          notification.Ts,
			MarkedState: notification.MarkedState,
		}

		if ntf.Data, err = applyUnitMeasure(ntf.AlertCode, ntf.Data); err != nil {
			continue
		}

		if notification.CustomerName != nil && *notification.CustomerName != "" {
			ntf.Subject = notification.CustomerName
		} else if notification.LicensePlate != nil && *notification.LicensePlate != "" {
			ntf.Subject = notification.LicensePlate
		}

		if notification.UrUsername != nil {
			usr := *notification.UrName + " " + *notification.UrSurname
			ntf.Subject = &usr
		}

		result = append(result, ntf)
	}

	return result, nil
}

// TODO da cambiare, in pratica deve dipendere dalla configurazione utente

func applyUnitMeasure(alertCode string, data any) (any, error) {
	j, ok := data.(datatypes.JSON)
	if !ok {
		return nil, fmt.Errorf("data is not datatype.JSON %T", data)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(j, &m); err != nil {
		return nil, err
	}

	v, ok := m["value"]
	if !ok {
		return nil, nil
	}

	str, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf(`"value" is not a string`)
	}

	if alertCode == enums.WarningPressure || alertCode == enums.AlertPressure {
		bar, err := formatValue(str)
		if err != nil {
			return nil, err
		}
		m["value"] = bar + " bar"
	}

	if alertCode == enums.WarningTemperature || alertCode == enums.AlertTemperature {
		m["value"] = str + " °C"
	}

	if strings.HasPrefix(alertCode, "TACHO") || strings.Contains(alertCode, "TIRE_CONTROL") {
		m["value"] = str + " km"
	}

	out, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return datatypes.JSON(out), nil
}

func formatValue(input string) (string, error) {
	// 1. parse
	v, err := strconv.Atoi(input)
	if err != nil {
		return "", err
	}

	// 2. convert to float
	f := float64(v) / 100.0

	// 3. format with rounding to 1 decimal
	return fmt.Sprintf("%.1f", f), nil
}
