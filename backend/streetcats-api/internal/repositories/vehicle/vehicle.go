package vehicle

import "sipli/notification-service/internal/entities"

type Repository interface {
	GetVehicleInfoByImei(imei string) (entities.VhVehicle, error)
	GetVehicleInfoByVehicleId(vehicleId uint64) (entities.VhVehicle, error)
}
