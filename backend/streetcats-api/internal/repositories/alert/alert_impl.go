package alert

import (
	"gorm.io/gorm"
)

type alertRepositoryImpl struct {
	pgCore *gorm.DB
}

func (n *alertRepositoryImpl) CreateAlertSchemas() error {
	//if err := n.pgCore.Exec("CREATE SCHEMA IF NOT EXISTS " + entities.AlertSchema).Error; err != nil {
	//	return err
	//}
	return nil
}

func NewAlertRepository(pgCore *gorm.DB) Repository {
	return &alertRepositoryImpl{pgCore: pgCore}
}
