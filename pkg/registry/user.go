package registry

import (
	"github.com/shabacha/pkg/usecase/usecase"

	"github.com/shabacha/pkg/adapter/controller"
	"github.com/shabacha/pkg/adapter/repository"
)

func (r *registry) NewUserController() controller.User {
	u := usecase.NewUserUsecase(
		repository.NewUserRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewUserController(u)
}
