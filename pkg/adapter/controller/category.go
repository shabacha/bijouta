package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/usecase"
)

type categoryController struct {
	categoryUsecase usecase.Category
}

type Category interface {
	GetCategories(c *gin.Context)
	GetCategory(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

func NewCategoryController(cs usecase.Category) Category {
	return &categoryController{cs}
}
func (cc *categoryController) GetCategories(ctx *gin.Context) {
	var u []*model.Category

	ca, err := cc.categoryUsecase.List(u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, ca)
}
func (cc *categoryController) CreateCategory(ctx *gin.Context) {
	var params model.Category

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ca, err := cc.categoryUsecase.Create(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, ca)
}

func (cc *categoryController) GetCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	ca, err := cc.categoryUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, ca)
	return
}
func (cc *categoryController) UpdateCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var params model.Category
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ca, err := cc.categoryUsecase.Update(&params, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, ca)
	return
}
func (cc *categoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := cc.categoryUsecase.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
	return
}
