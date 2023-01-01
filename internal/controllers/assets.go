package controllers

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Asset struct {
	*Base
}

func (c *Asset) GetAllAssets(context *gin.Context) {
	assets := map[string]interface{}{}
	db.DB.Model(&models.Asset{}).Find(assets)

	context.JSON(http.StatusOK, api.ResFromData(assets))
}

func (c *Asset) CreateAsset(context *gin.Context) {
	newAsset := &models.Asset{}

	err := context.BindJSON(newAsset)
	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))
		return
	}

	db.DB.Create(newAsset)
	context.JSON(http.StatusCreated, api.ResFromData(newAsset))
}
