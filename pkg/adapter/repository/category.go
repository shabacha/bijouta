package repository

import (
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{db}
}
func (cr *categoryRepository) FindAll(c []*model.Category) ([]*model.Category, error) {
	err := cr.db.Find(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}
func (cr *categoryRepository) Create(c *model.Category) (*model.Category, error) {
	// Call Create() on the database object to store the user
	if err := cr.db.Create(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
func (cr *categoryRepository) GetById(id int) (*model.Category, error) {
	var c model.Category
	cr.db.Where("id = ?", id).First(&c)
	return &c, nil
}
func (cr *categoryRepository) Update(c *model.Category, id int) (*model.Category, error) {
	if err := cr.db.Model(&model.Category{}).Where("id = ?", id).Updates(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
func (cr *categoryRepository) Delete(id int) error {
	if err := cr.db.Model(&model.Category{}).Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		return err
	}
	return nil
}
