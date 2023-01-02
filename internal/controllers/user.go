package controllers

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/internal/utils"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Base
}

func (u *User) CreateUser(context *gin.Context) {
	newUser := &models.User{}

	err := context.BindJSON(newUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))
		return
	}

	if !utils.ValidatePassword(newUser.Password) {
		context.JSON(http.StatusBadRequest, api.ResFromMessage("invalid password prompted"))
		return
	}

	newPass, err := utils.HashPassword(newUser.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))

		return
	}

	newUser.Password = newPass
	err = db.DB.Create(newUser).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))

		return
	}

	context.JSON(http.StatusCreated, api.Response{Data: newUser})
}

func (u *User) GetAllUser(context *gin.Context) {
	var users []models.User

	err := db.DB.Find(&users).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, api.ResFromError(err))
		return
	}

	context.JSON(http.StatusOK, api.Response{Data: users})
}
