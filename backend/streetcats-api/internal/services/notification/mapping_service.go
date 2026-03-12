package notification

import (
	"encoding/json"
	"errors"
	"fmt"
	"sipli/notification-service/internal/dto"
	"sipli/notification-service/internal/entities"
	"sipli/notification-service/internal/enums"
	"strconv"
	"strings"

	"gorm.io/datatypes"
)

func (s *Service) MappingMessage(message dto.NotificationEventDTO) (dto.NotificationDTO, error) {
	var (
		vhInfo      entities.VhVehicle
		companyName string
		id          uint64
		err         error
	)

	messageToSent := dto.NotificationDTO{
		ID:          message.ID,
		Category:    message.Category,
		Severity:    message.Severity,
		AlertCode:   message.AlertCode,
		Data:        message.Data,
		Ts:          message.Ts,
		MarkedState: false,
	}

	if messageToSent.Data != nil {
		if messageToSent.Data, err = applyUnitMeasureOnStream(messageToSent.AlertCode, messageToSent.Data); err != nil {
			return messageToSent, err
		}
	}

	switch message.SourceType {
	case enums.Imei:
		if vhInfo, err = s.vehicleRepo.GetVehicleInfoByImei(message.Source); err != nil {
			return messageToSent, err
		}
		messageToSent.Subject = &vhInfo.LicensePlate
	case enums.VehicleId:
		if id, err = strconv.ParseUint(message.Source, 10, 64); err == nil {
			return messageToSent, nil
		}

		if vhInfo, err = s.vehicleRepo.GetVehicleInfoByVehicleId(id); err != nil {
			return messageToSent, err
		}

		messageToSent.Subject = &vhInfo.LicensePlate
	case enums.CustomerId:
		if id, err = strconv.ParseUint(message.Source, 10, 64); err != nil {
			return messageToSent, err
		}
		if companyName, err = s.customerRepo.GetCustomerInfoById(id); err != nil {
			return messageToSent, err
		}
		messageToSent.Subject = &companyName
	default:
		return messageToSent, errors.New("")

	}

	return messageToSent, nil
}

func applyUnitMeasureOnStream(alertCode string, data any) (any, error) {
	m, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("data is not datatype.JSON %T", data)
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
