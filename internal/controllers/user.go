package controllers

import (
	"github.com/MatteoMiotello/goAccounting/internal/utils"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	var newUser models.User
	newUser.Email = userDto.Email
	newUser.Password = newPass
	err = newUser.InsertG(context, boil.Infer())

	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))

		return
	}

	context.JSON(http.StatusCreated, api.Response{Data: newUser})
}

func (u *User) GetAllUser(context *gin.Context) {

	users, err := models.Users().AllG(context)

	if err != nil {
		context.JSON(http.StatusInternalServerError, api.ResFromError(err))
		return
	}

	context.JSON(http.StatusOK, api.Response{Data: users})
}
