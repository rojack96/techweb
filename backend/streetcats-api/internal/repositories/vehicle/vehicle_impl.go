package vehicle

import (
	"sipli/notification-service/internal/entities"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type vehicleRepositoryImpl struct {
	pgCore *gorm.DB
	log    *zap.Logger
}

func (n *vehicleRepositoryImpl) GetVehicleInfoByImei(imei string) (entities.VhVehicle, error) {
	var result entities.VhVehicle

	if err := n.pgCore.Select("vh.id, vh.licenseplate").
		Table("crmdb.ws_units wu").
		Joins("JOIN crmdb.dv_devices dv ON dv.idunit = wu.id").
		Joins("JOIN crmdb.vh_vehicle vh ON vh.id = dv.idvehicle").
		Where("wu.imei = ?", imei).
		Take(&result).Error; err != nil {
		return entities.VhVehicle{}, err
	}

	return result, nil
}

func (n *vehicleRepositoryImpl) GetVehicleInfoByVehicleId(vehicleId uint64) (entities.VhVehicle, error) {
	var result entities.VhVehicle

	if err := n.pgCore.Where("vh.id = ?", vehicleId).
		First(&result).Error; err != nil {
		return entities.VhVehicle{}, err
	}

	return result, nil
}

func NewVehicleRepository(pgCore *gorm.DB, log *zap.Logger) Repository {
	return &vehicleRepositoryImpl{pgCore: pgCore, log: log}
}
