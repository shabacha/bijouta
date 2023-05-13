package repository

import "github.com/shabacha/pkg/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
	GetById(id int) (*model.User, error)
	Update(u *model.User, id int) (*model.User, error)
	Delete(id int) error
	GetUserByUserName(u *model.User, username string) (*model.User, error)
}
