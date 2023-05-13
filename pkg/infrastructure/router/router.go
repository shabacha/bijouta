package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shabacha/pkg/adapter/controller"
	util "github.com/shabacha/pkg/util/jwt"
)

func NewRouter(r *gin.Engine, c controller.AppController) *gin.Engine {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	userGroup := r.Group("/users", util.TokenAuthMiddleware())
	userGroup.GET("", c.User.GetUsers)
	userGroup.GET("/:id", c.User.GetUser)
	userGroup.PUT("/:id", c.User.UpdateUser)
	userGroup.DELETE("/:id", c.User.DeleteUser)
	authGroup := r.Group("/auth")
	authGroup.POST("/login", c.User.Login)
	authGroup.POST("/register", c.User.CreateUser)
	// productGroup := r.Group("/products")
	// productGroup.POST("", func(context echo.Context) error { return c.Product.CreateProduct(context) })
	// productGroup.GET("", func(context echo.Context) error { return c.Product.GetAllProducts(context) })
	// productGroup.GET("/:id", func(context echo.Context) error { return c.Product.GetProduct(context) })
	// productGroup.PUT("/:id", func(context echo.Context) error { return c.Product.UpdateProduct(context) })
	// productGroup.DELETE("/:id", func(context echo.Context) error { return c.Product.DeleteProduct(context) })

	return r
}
