package controllers

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/internal/utils"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/MatteoMiotello/goAccounting/pkg/security"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type Login struct {
	Base
}

type SignInDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *Login) SignIn(ctx *gin.Context) {
	signInDto := SignInDto{}

	err := ctx.BindJSON(&signInDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	var user models.User

	err = db.DB.Find(&user, "email = ?", signInDto.Email).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{Error: "incorrect user or password"})
		return
	}

	if !utils.CheckPassword(signInDto.Password, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.ErrorResponse{Error: "incorrect user or password"})
		return
	}

	var signedToken string
	signedToken, err = security.BuildJwtFromUser(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.SetCookie(viper.GetString("security.jwt.cookie-key"), signedToken, viper.GetInt("security.jwt.expiration"), "/", "localhost", true, true)
	ctx.JSON(http.StatusOK, api.Response{Data: user})
}
