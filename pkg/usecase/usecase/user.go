package usecase

import (
	"errors"

	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
	dBRepository   repository.DBRepository
}

type User interface {
	List(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
	Get(id int) (*model.User, error)
	Update(u *model.User, id int) (*model.User, error)
	Login(infos *model.LoginInput) (*model.User, error)
}

func NewUserUsecase(r repository.UserRepository, d repository.DBRepository) User {
	return &userUsecase{r, d}
}

func (uu *userUsecase) List(u []*model.User) ([]*model.User, error) {
	u, err := uu.userRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) Create(u *model.User) (*model.User, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.userRepository.Create(u)
		if err != nil {
			return nil, err
		}
		return u, nil
	})
	user, ok := data.(*model.User)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("Error casting user")
	}

	return user, nil
}

func (uu *userUsecase) Get(id int) (*model.User, error) {
	u, err := uu.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uu *userUsecase) Update(u *model.User, id int) (*model.User, error) {
	u, err := uu.userRepository.Update(u, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (uu *userUsecase) Login(infos *model.LoginInput) (*model.User, error) {
	return nil, nil
}
