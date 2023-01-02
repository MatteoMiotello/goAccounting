package controllers

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Asset struct {
	Base
}

type createAssetDto struct {
	Name   string `json:"name" binding:"required"`
	UserId uint   `json:"userId" binding:"required"`
}

func (c *Asset) GetAllAssets(context *gin.Context) {
	var assets []models.Asset
	err := db.DB.Preload("User").Find(&assets).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, api.ResFromError(err))
		return
	}

	context.JSON(http.StatusOK, api.Response{
		Data: assets,
	})
}

func (c *Asset) CreateAsset(context *gin.Context) {
	assetData := &createAssetDto{}
	err := context.BindJSON(assetData)

	if err != nil {
		context.JSON(http.StatusBadRequest, api.ResFromError(err))
		return
	}

	err = db.DB.First(&models.User{}, assetData.UserId).Where("deleted_at is null").Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, api.ResFromError(err))
		return
	}

	newAsset := &models.Asset{Name: assetData.Name, UserID: assetData.UserId}

	err = db.DB.Create(newAsset).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, api.ResFromError(err))
		return
	}

	context.JSON(http.StatusCreated, api.Response{Data: newAsset})
}
