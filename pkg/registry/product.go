package registry

import (
	"github.com/shabacha/pkg/adapter/controller"
	"github.com/shabacha/pkg/adapter/repository"
	"github.com/shabacha/pkg/usecase/usecase"
)

func (r *registry) NewProductController() controller.Product {
	c := usecase.NewProductUsecase(
		repository.NewProductRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewProductController(c)
}
