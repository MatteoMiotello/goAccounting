package bootstrap

import (
	"github.com/MatteoMiotello/goAccounting/internal/controllers"
	"github.com/gin-gonic/gin"
)

type Api struct {
}

func registerRoutes(router *gin.Engine) {
	userGroup := router.Group("user")
	userController := new(controllers.User)
	userGroup.GET("", userController.GetAllUser)
	userGroup.POST("", userController.CreateUser)

	assetGroup := router.Group("asset")
	assetController := new(controllers.Asset)
	assetGroup.GET("", assetController.GetAllAssets)
	assetGroup.POST("", assetController.CreateAsset)
}

func (a Api) Init() error {
	router := gin.Default()

	registerRoutes(router)

	err := router.Run()
	if err != nil {
		return err
	}

	return nil
}
