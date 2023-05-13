package repository

import "github.com/shabacha/pkg/domain/model"

type ProductRepository interface {
	FindAll(u []*model.Product) ([]*model.Product, error)
	Create(u *model.Product) (*model.Product, error)
	GetById(id int) (*model.Product, error)
	Update(u *model.Product, id int) (*model.Product, error)
	Delete(id int) error
}
