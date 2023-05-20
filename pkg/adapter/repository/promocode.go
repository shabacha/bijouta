package repository

import (
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
	"gorm.io/gorm"
)

type promocodeRepository struct {
	db *gorm.DB
}

func NewPromoCodeRepository(db *gorm.DB) repository.PromoCodeRepository {
	return &promocodeRepository{db}
}
func (pcr *promocodeRepository) FindAll(c []*model.PromoCode) ([]*model.PromoCode, error) {
	err := pcr.db.Find(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}
func (pcr *promocodeRepository) Create(c *model.PromoCode) (*model.PromoCode, error) {
	// Call Create() on the database object to store the user
	if err := pcr.db.Create(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
func (pcr *promocodeRepository) GetById(id int) (*model.PromoCode, error) {
	var pc model.PromoCode
	pcr.db.Where("id = ?", id).First(&pc)
	return &pc, nil
}
func (pcr *promocodeRepository) Update(c *model.PromoCode, id int) (*model.PromoCode, error) {
	if err := pcr.db.Model(&model.PromoCode{}).Where("id = ?", id).Updates(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
func (pcr *promocodeRepository) Delete(id int) error {
	if err := pcr.db.Model(&model.PromoCode{}).Where("id = ?", id).Delete(&model.PromoCode{}).Error; err != nil {
		return err
	}
	return nil
}
