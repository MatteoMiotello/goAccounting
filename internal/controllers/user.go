package controllers

import (
	"github.com/MatteoMiotello/goAccounting/internal/repo"
	"github.com/MatteoMiotello/goAccounting/internal/utils"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User Base

type createUserDto struct {
	Email    string
	Password string
}

func (u *User) CreateUser(context *gin.Context) {
	userDto := new(createUserDto)

	err := context.BindJSON(userDto)

	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))
		return
	}

	if !utils.ValidatePassword(userDto.Password) {
		context.JSON(http.StatusBadRequest, api.ResFromMessage("invalid password prompted"))
		return
	}

	newPass, err := utils.HashPassword(userDto.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))

		return
	}

	userRepo := repo.NewUserRepo()

	var newUser models.User
	newUser.Email = userDto.Email
	newUser.Password = newPass

	err = userRepo.Insert(context, &newUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))

		return
	}

	context.JSON(http.StatusCreated, api.Response{Data: newUser})
}

func (u *User) GetAllUser(context *gin.Context) {
	userRepo := repo.NewUserRepo()
	users, err := userRepo.GetAll(context)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, api.ResFromError(err))
		return
	}

	context.JSON(http.StatusOK, api.Response{Data: users})
}
