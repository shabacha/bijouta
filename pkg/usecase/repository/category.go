package repository

import "github.com/shabacha/pkg/domain/model"

type CategoryRepository interface {
	FindAll(u []*model.Category) ([]*model.Category, error)
	Create(u *model.Category) (*model.Category, error)
	GetById(id int) (*model.Category, error)
	Update(u *model.Category, id int) (*model.Category, error)
	Delete(id int) error
}
