package usecase

import (
	"errors"

	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
)

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
	dBRepository       repository.DBRepository
}
type Category interface {
	List(u []*model.Category) ([]*model.Category, error)
	Create(u *model.Category) (*model.Category, error)
	Get(id int) (*model.Category, error)
	Update(u *model.Category, id int) (*model.Category, error)
	Delete(id int) error
}

func NewCategoryUsecase(r repository.CategoryRepository, d repository.DBRepository) Category {
	return &categoryUsecase{r, d}
}
func (cu *categoryUsecase) List(u []*model.Category) ([]*model.Category, error) {
	c, err := cu.categoryRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return c, nil
}
func (cu *categoryUsecase) Create(u *model.Category) (*model.Category, error) {
	data, err := cu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		c, err := cu.categoryRepository.Create(u)
		if err != nil {
			return nil, err
		}
		return c, nil
	})
	category, ok := data.(*model.Category)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("invalid data schema for category")
	}

	return category, nil
}
func (cu *categoryUsecase) Get(id int) (*model.Category, error) {
	c, err := cu.categoryRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cu *categoryUsecase) Update(u *model.Category, id int) (*model.Category, error) {
	c, err := cu.categoryRepository.Update(u, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}
func (cu *categoryUsecase) Delete(id int) error {
	if err := cu.categoryRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
