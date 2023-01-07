package controllers

import (
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Transaction Base

type transactionCategoryDto struct {
	Name string `json:"name" binding:"required"`
}

type transactionDto struct {
	AssetId     uint                   `json:"assetId" binding:"required"`
	Category    transactionCategoryDto `json:"category" binding:"required"`
	Amount      float64                `json:"amount" binding:"required"`
	Description string                 `json:"description" binding:"required"`
}

func (t *Transaction) CreateTransaction(ctx *gin.Context) {
	var tDto transactionDto
	var err error
	err = ctx.BindJSON(&tDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	var tCategory models.TransactionCategory
	count := new(int64)
	db.DB.First(&tCategory, "name = ?", tDto.Category.Name).Count(count)

	tModel := models.Transaction{
		AssetID:     tDto.AssetId,
		Description: tDto.Description,
		Amount:      tDto.Amount,
	}

	if *count == 0 {
		tModel.TransactionCategory = models.TransactionCategory{Name: tDto.Category.Name}
	} else {
		tModel.TransactionCategoryID = tCategory.ID
	}

	err = db.DB.Create(&tModel).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tModel)
}
