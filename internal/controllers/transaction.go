package controllers

import (
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
	"net/http"
)

type Transaction Base

type transactionCategoryDto struct {
	Name string `json:"name" binding:"required"`
}

type transactionDto struct {
	AssetId     int64                  `json:"assetId" binding:"required"`
	Category    transactionCategoryDto `json:"category" binding:"required"`
	Amount      types.Decimal          `json:"amount" binding:"required"`
	Description null.String            `json:"description" binding:"required"`
}

func (t *Transaction) CreateTransaction(ctx *gin.Context) {
	var tDto transactionDto
	var err error
	err = ctx.BindJSON(&tDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
		return
	}

	tModel := models.Transaction{
		AssetID:     tDto.AssetId,
		Description: tDto.Description,
		Amount:      tDto.Amount,
	}
	tCategory, err := models.TransactionCategories(qm.Where("name=?", tDto.Category.Name)).OneG(ctx)

	if err == nil {
		tCategory := &models.TransactionCategory{Name: tDto.Category.Name}
		tModel.SetTransactionCategoryG(ctx, true, tCategory)
	} else {
		tModel.TransactionCategoryID = tCategory.ID
	}

	err = tModel.InsertG(ctx, boil.Infer())

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tModel)
}
