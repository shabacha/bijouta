package repository

import (
	"log"

	"github.com/shabacha/pkg/usecase/repository"

	"gorm.io/gorm"
)

type dbRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) repository.DBRepository {
	return &dbRepository{db}
}

func (r *dbRepository) Transaction(txFunc func(interface{}) (interface{}, error)) (data interface{}, err error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			log.Print("recover")
			tx.Rollback()
		} else if err != nil {
			log.Print("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	data, err = txFunc(tx)
	return data, err
}
