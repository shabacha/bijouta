package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shabacha/pkg/usecase/usecase"

	"github.com/shabacha/pkg/domain/model"
)

type userController struct {
	userUsecase usecase.User
}

type User interface {
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

func NewUserController(us usecase.User) User {
	return &userController{us}
}

func (uc *userController) GetUsers(ctx *gin.Context) {
	var u []*model.User

	u, err := uc.userUsecase.List(u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	var params model.User

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	u, err := uc.userUsecase.Create(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, u)
}

func (uc *userController) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	u, err := uc.userUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, u)

}
func (uc *userController) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var params model.User

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	u, err := uc.userUsecase.Update(&params, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, u)
	return
}
