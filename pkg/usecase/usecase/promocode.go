package usecase

import (
	"errors"

	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
)

type promocodeUsecase struct {
	promocodeRepository repository.PromoCodeRepository
	dBRepository        repository.DBRepository
}

type PromoCode interface {
	List(u []*model.PromoCode) ([]*model.PromoCode, error)
	Create(u *model.PromoCode) (*model.PromoCode, error)
	Get(id int) (*model.PromoCode, error)
	Update(u *model.PromoCode, id int) (*model.PromoCode, error)
	Delete(id int) error
}

func NewPromoCodeUsecase(r repository.PromoCodeRepository, d repository.DBRepository) PromoCode {
	return &promocodeUsecase{r, d}
}
func (pcu *promocodeUsecase) List(pcs []*model.PromoCode) ([]*model.PromoCode, error) {
	pc, err := pcu.promocodeRepository.FindAll(pcs)
	if err != nil {
		return nil, err
	}

	return pc, nil
}
func (pcu *promocodeUsecase) Create(u *model.PromoCode) (*model.PromoCode, error) {
	data, err := pcu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		pc, err := pcu.promocodeRepository.Create(u)
		if err != nil {
			return nil, err
		}
		return pc, nil
	})
	promocode, ok := data.(*model.PromoCode)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("invalid data schema for category")
	}

	return promocode, nil
}
func (pcu *promocodeUsecase) Get(id int) (*model.PromoCode, error) {
	pc, err := pcu.promocodeRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return pc, nil
}

func (pcu *promocodeUsecase) Update(pcm *model.PromoCode, id int) (*model.PromoCode, error) {
	pc, err := pcu.promocodeRepository.Update(pcm, id)
	if err != nil {
		return nil, err
	}
	return pc, nil
}
func (pcu *promocodeUsecase) Delete(id int) error {
	if err := pcu.promocodeRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
