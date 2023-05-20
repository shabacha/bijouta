package controller

type AppController struct {
	User      interface{ User }
	Category  interface{ Category }
	Product   interface{ Product }
	PromoCode interface{ PromoCode }
}
