package usecase

import (
	"errors"

	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
)

type productUsecase struct {
	productRepository repository.ProductRepository
	dBRepository      repository.DBRepository
}

type Product interface {
	List(u []*model.Product) ([]*model.Product, error)
	Create(u *model.Product) (*model.Product, error)
	Get(id int) (*model.Product, error)
	Update(u *model.Product, id int) (*model.Product, error)
	Delete(id int) error
}

func NewProductUsecase(r repository.ProductRepository, d repository.DBRepository) Product {
	return &productUsecase{r, d}
}
func (pu *productUsecase) List(u []*model.Product) ([]*model.Product, error) {
	p, err := pu.productRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return p, nil
}
func (pu *productUsecase) Create(u *model.Product) (*model.Product, error) {
	data, err := pu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		p, err := pu.productRepository.Create(u)
		if err != nil {
			return nil, err
		}
		return p, nil
	})
	product, ok := data.(*model.Product)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("invalid data schema for category")
	}

	return product, nil
}
func (pu *productUsecase) Get(id int) (*model.Product, error) {
	p, err := pu.productRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pu *productUsecase) Update(pm *model.Product, id int) (*model.Product, error) {
	p, err := pu.productRepository.Update(pm, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
func (pu *productUsecase) Delete(id int) error {
	if err := pu.productRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
