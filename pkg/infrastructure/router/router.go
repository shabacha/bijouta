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
	categoryGroup := r.Group("/categories")
	categoryGroup.GET("", c.Category.GetCategories)
	categoryGroup.GET("/:id", c.Category.GetCategory)
	categoryGroup.PUT("/:id", c.Category.UpdateCategory)
	categoryGroup.DELETE("/:id", c.Category.DeleteCategory)
	categoryGroup.POST("/create", c.Category.CreateCategory)
	productGroup := r.Group("/products")
	productGroup.POST("", c.Product.CreateProduct)
	productGroup.GET("", c.Product.GetProducts)
	productGroup.GET("/:id", c.Product.GetProduct)
	productGroup.PUT("/:id", c.Product.UpdateProduct)
	productGroup.DELETE("/:id", c.Product.DeleteProduct)

	return r
}
