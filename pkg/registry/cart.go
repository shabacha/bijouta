package registry

import (
	"github.com/shabacha/pkg/adapter/controller"
	"github.com/shabacha/pkg/adapter/repository"
	"github.com/shabacha/pkg/usecase/usecase"
)

func (r *registry) NewCartController() controller.Cart {
	c := usecase.NewCartUsecase(
		repository.NewCartRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewCartController(c)
}
