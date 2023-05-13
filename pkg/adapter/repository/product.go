package repository

import (
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db}
}
func (pr *productRepository) FindAll(c []*model.Product) ([]*model.Product, error) {
	err := pr.db.Find(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}
func (pr *productRepository) Create(c *model.Product) (*model.Product, error) {
	// Call Create() on the database object to store the user
	if err := pr.db.Create(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
func (pr *productRepository) GetById(id int) (*model.Product, error) {
	var c model.Product
	pr.db.Where("id = ?", id).First(&c)
	return &c, nil
}
func (pr *productRepository) Update(c *model.Product, id int) (*model.Product, error) {
	if err := pr.db.Model(&model.Product{}).Where("id = ?", id).Updates(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
func (pr *productRepository) Delete(id int) error {
	if err := pr.db.Model(&model.Product{}).Where("id = ?", id).Delete(&model.Product{}).Error; err != nil {
		return err
	}
	return nil
}
