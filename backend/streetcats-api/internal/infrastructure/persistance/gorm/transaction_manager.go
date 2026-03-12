package gorm

import "gorm.io/gorm"

type TransactionManager interface {
	Transaction(fn func(tx any) error) error
}

type gormTransactionManager struct {
	db *gorm.DB
}

func NewGormTransactionManager(db *gorm.DB) TransactionManager {
	return &gormTransactionManager{db: db}
}

func (t *gormTransactionManager) Transaction(fn func(tx any) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
