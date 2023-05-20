package usecase

import (
	"errors"

	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
)

type cartUsecase struct {
	cartRepository repository.CartRepository
	dBRepository   repository.DBRepository
}
type Cart interface {
	List(u []*model.Cart) ([]*model.Cart, error)
	Create(u *model.Cart) (*model.Cart, error)
	Get(id int) (*model.Cart, error)
	Update(u *model.Cart, id int) (*model.Cart, error)
	Delete(id int) error
}

func NewCartUsecase(r repository.CartRepository, d repository.DBRepository) Cart {
	return &cartUsecase{r, d}
}
func (ctu *cartUsecase) List(u []*model.Cart) ([]*model.Cart, error) {
	c, err := ctu.cartRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return c, nil
}
func (ctu *cartUsecase) Create(u *model.Cart) (*model.Cart, error) {
	data, err := ctu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		c, err := ctu.cartRepository.Create(u)
		if err != nil {
			return nil, err
		}
		return c, nil
	})
	Cart, ok := data.(*model.Cart)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("invalid data schema for Cart")
	}

	return Cart, nil
}
func (ctu *cartUsecase) Get(id int) (*model.Cart, error) {
	c, err := ctu.cartRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (ctu *cartUsecase) Update(u *model.Cart, id int) (*model.Cart, error) {
	c, err := ctu.cartRepository.Update(u, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}
func (ctu *cartUsecase) Delete(id int) error {
	if err := ctu.cartRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
