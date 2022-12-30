package api

import (
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AssetsRepository struct {
	Repository
}

func getAllAssets(context *gin.Context) {
	assets := map[string]interface{}{}
	db.DB.Model(&models.Asset{}).Find(assets)

	context.JSON(http.StatusOK, gin.H{"data": assets})
}

func createAsset(context *gin.Context) {
	newAsset := &models.Asset{}

	err := context.BindJSON(newAsset)
	if err != nil {
		return
	}

	db.DB.Create(newAsset)

	context.JSON(http.StatusCreated, gin.H{"data": newAsset})
}

func (repo AssetsRepository) HandleRequests() error {
	group := repo.Api.Group("/assets")

	group.GET("", getAllAssets)
	group.POST("", createAsset)

	return nil
}
