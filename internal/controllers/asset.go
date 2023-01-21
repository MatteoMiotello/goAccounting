package controllers

import (
	"github.com/MatteoMiotello/goAccounting/internal/repo"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Asset Base

type createAssetDto struct {
	Name   string `json:"name" binding:"required"`
	UserId int64  `json:"userId" binding:"required"`
}

func (c *Asset) GetAllAssets(context *gin.Context) {
	assetRepo := new(repo.AssetRepo)

	assets, err := assetRepo.GetAll(context)

	if err != nil {
		return
	}

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

	var newAsset models.Asset
	newAsset.Name = assetData.Name
	newAsset.UserID = assetData.UserId

	assetRepo := repo.NewAssetRepo()
	err = assetRepo.Insert(context, &newAsset)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, api.ResFromError(err))
		return
	}

	context.JSON(http.StatusCreated, api.Response{Data: newAsset})
}
