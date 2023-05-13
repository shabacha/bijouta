package registry

import (
	"github.com/shabacha/pkg/adapter/controller"
	"github.com/shabacha/pkg/adapter/repository"
	"github.com/shabacha/pkg/usecase/usecase"
)

func (r *registry) NewCategoryController() controller.Category {
	c := usecase.NewCategoryUsecase(
		repository.NewCategoryRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewCategoryController(c)
}
