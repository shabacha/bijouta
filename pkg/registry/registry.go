package registry

import (
	"github.com/shabacha/pkg/adapter/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User:      r.NewUserController(),
		Category:  r.NewCategoryController(),
		Product:   r.NewProductController(),
		PromoCode: r.NewPromoCodeController(),
		Cart:      r.NewCartController(),
	}
}
