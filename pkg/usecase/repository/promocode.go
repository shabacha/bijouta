package repository

import "github.com/shabacha/pkg/domain/model"

type PromoCodeRepository interface {
	FindAll(u []*model.PromoCode) ([]*model.PromoCode, error)
	Create(u *model.PromoCode) (*model.PromoCode, error)
	GetById(id int) (*model.PromoCode, error)
	Update(u *model.PromoCode, id int) (*model.PromoCode, error)
	Delete(id int) error
}
