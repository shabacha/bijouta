package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/usecase"
)

type productController struct {
	productUsecase usecase.Product
}

type Product interface {
	GetProducts(c *gin.Context)
	GetProduct(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

func NewProductController(ps usecase.Product) Product {
	return &productController{ps}
}
func (pc *productController) GetProducts(ctx *gin.Context) {
	var pu []*model.Product

	p, err := pc.productUsecase.List(pu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, p)
}
func (pc *productController) CreateProduct(ctx *gin.Context) {
	var params model.Product

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	p, err := pc.productUsecase.Create(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, p)
}

func (pc *productController) GetProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	p, err := pc.productUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, p)
	return
}
func (pc *productController) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var params model.Product
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	p, err := pc.productUsecase.Update(&params, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, p)
	return
}
func (pc *productController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := pc.productUsecase.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	return
}
