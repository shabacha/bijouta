package util

import "github.com/shabacha/pkg/domain/model"

func ApplyPromoCode(product *model.Product, promoCode *model.PromoCode) {
	discountedPrice := product.Price * (1 - promoCode.Discount)
	product.Price = discountedPrice
}
