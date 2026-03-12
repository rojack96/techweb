package customer

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type customerRepositoryImpl struct {
	pgCore *gorm.DB
	log    *zap.Logger
}

func (n *customerRepositoryImpl) GetCustomerInfoById(id uint64) (string, error) {
	var result string

	if err := n.pgCore.Select("cl.companyname").
		Table("crmdb.cl_clients AS cl").Where("cl.id = ?", id).
		Take(&result).Error; err != nil {
		return "", err
	}

	return result, nil
}

func NewCustomerRepository(pgCore *gorm.DB, log *zap.Logger) Repository {
	return &customerRepositoryImpl{pgCore: pgCore, log: log}
}
