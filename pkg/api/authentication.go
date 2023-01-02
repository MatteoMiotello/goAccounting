package api

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/security"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func IsAuthenticated(ctx *gin.Context) {
	var err error
	token, err := ctx.Cookie(viper.GetString("security.jwt.cookie-key"))
	if err != nil {
		return
	}
	var userClaims *security.UserJwtClaims
	userClaims, err = security.ParseUserJwt(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var userModel models.User
	err = db.DB.First(&userModel, userClaims.UserId).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Error: "user not found",
		})
		return
	}

	ctx.Set("user", &userModel)
	ctx.Next()
}
