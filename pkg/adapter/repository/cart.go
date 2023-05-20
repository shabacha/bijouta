package repository

import (
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) repository.CartRepository {
	return &cartRepository{db}
}
func (ctr *cartRepository) FindAll(c []*model.Cart) ([]*model.Cart, error) {
	err := ctr.db.Find(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}
func (ctr *cartRepository) Create(c *model.Cart) (*model.Cart, error) {
	// Call Create() on the database object to store the user
	if err := ctr.db.Create(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
func (ctr *cartRepository) GetById(id int) (*model.Cart, error) {
	var c model.Cart
	ctr.db.Where("id = ?", id).First(&c)
	return &c, nil
}
func (ctr *cartRepository) Update(c *model.Cart, id int) (*model.Cart, error) {
	if err := ctr.db.Model(&model.Cart{}).Where("id = ?", id).Updates(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
func (ctr *cartRepository) Delete(id int) error {
	if err := ctr.db.Model(&model.Cart{}).Where("id = ?", id).Delete(&model.Cart{}).Error; err != nil {
		return err
	}
	return nil
}
