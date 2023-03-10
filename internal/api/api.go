package api

import (
	"github.com/MatteoMiotello/goAccounting/internal/controllers"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	loginGroup := router.Group("login")
	loginController := new(controllers.Login)
	loginGroup.POST("", loginController.SignIn)

	userGroup := router.Group("user")
	userGroup.Use(api.IsAuthenticated)
	userController := new(controllers.User)
	userGroup.GET("", userController.GetAllUser)
	userGroup.POST("", userController.CreateUser)

	assetGroup := router.Group("asset")
	assetGroup.Use(api.IsAuthenticated)
	assetController := new(controllers.Asset)
	assetGroup.GET("", assetController.GetAllAssets)
	assetGroup.POST("", assetController.CreateAsset)

	transactionGroup := router.Group("transaction")
	transactionGroup.Use(api.IsAuthenticated)
	transactionController := new(controllers.Transaction)
	transactionGroup.POST("", transactionController.CreateTransaction)
}

func Run() {
	router := gin.Default()

	registerRoutes(router)

	err := router.Run()
	if err != nil {
		panic("error initializing api routes: " + err.Error())
	}
}
