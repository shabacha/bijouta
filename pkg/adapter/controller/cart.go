package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/usecase"
)

type cartController struct {
	cartUsecase usecase.Cart
}

type Cart interface {
	GetCarts(c *gin.Context)
	GetCart(c *gin.Context)
	CreateCart(c *gin.Context)
	UpdateCart(c *gin.Context)
	DeleteCart(c *gin.Context)
}

func NewCartController(cs usecase.Cart) Cart {
	return &cartController{cs}
}
func (ctc *cartController) GetCarts(ctx *gin.Context) {
	var u []*model.Cart

	ca, err := ctc.cartUsecase.List(u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, ca)
}
func (ctc *cartController) CreateCart(ctx *gin.Context) {
	var params model.Cart

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ca, err := ctc.cartUsecase.Create(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, ca)
}

func (ctc *cartController) GetCart(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	ca, err := ctc.cartUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, ca)
	return
}
func (ctc *cartController) UpdateCart(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var params model.Cart
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ca, err := ctc.cartUsecase.Update(&params, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, ca)
	return
}
func (ctc *cartController) DeleteCart(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctc.cartUsecase.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
	return
}
