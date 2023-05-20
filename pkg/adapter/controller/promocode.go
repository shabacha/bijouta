package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/usecase"
)

type promocodeController struct {
	promocodeUsecase usecase.PromoCode
}

type PromoCode interface {
	GetPromoCodes(c *gin.Context)
	GetPromoCode(c *gin.Context)
	CreatePromoCode(c *gin.Context)
	UpdatePromoCode(c *gin.Context)
	DeletePromoCode(c *gin.Context)
}

func NewPromoCodeController(ps usecase.PromoCode) PromoCode {
	return &promocodeController{ps}
}
func (pc *promocodeController) GetPromoCodes(ctx *gin.Context) {
	var pcs []*model.PromoCode

	pcl, err := pc.promocodeUsecase.List(pcs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, pcl)
}
func (pc *promocodeController) CreatePromoCode(ctx *gin.Context) {
	var params model.PromoCode

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ca, err := pc.promocodeUsecase.Create(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, ca)
}

func (pc *promocodeController) GetPromoCode(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	ca, err := pc.promocodeUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, ca)
	return
}
func (pc *promocodeController) UpdatePromoCode(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var params model.PromoCode
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ca, err := pc.promocodeUsecase.Update(&params, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, ca)
	return
}
func (pc *promocodeController) DeletePromoCode(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := pc.promocodeUsecase.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Promocode deleted successfully"})
	return
}
