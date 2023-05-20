package registry

import (
	"github.com/shabacha/pkg/adapter/controller"
	"github.com/shabacha/pkg/adapter/repository"
	"github.com/shabacha/pkg/usecase/usecase"
)

func (r *registry) NewPromoCodeController() controller.PromoCode {
	pc := usecase.NewPromoCodeUsecase(
		repository.NewPromoCodeRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewPromoCodeController(pc)
}
