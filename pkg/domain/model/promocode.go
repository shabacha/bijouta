package model

type PromoCode struct {
	Code     string
	Discount float64
	Valid    bool
}

func (PromoCode) TableName() string { return "promo_codes" }
