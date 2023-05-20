package repository

import "github.com/shabacha/pkg/domain/model"

type CartRepository interface {
	FindAll(u []*model.Cart) ([]*model.Cart, error)
	Create(u *model.Cart) (*model.Cart, error)
	GetById(id int) (*model.Cart, error)
	Update(u *model.Cart, id int) (*model.Cart, error)
	Delete(id int) error
}
