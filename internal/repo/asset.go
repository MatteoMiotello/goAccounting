package repo

import (
	"context"
	"github.com/MatteoMiotello/goAccounting/internal/db"
	"github.com/MatteoMiotello/goAccounting/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AssetRepo BaseRepo

func NewAssetRepo() *AssetRepo {
	return new(AssetRepo)
}

func (a AssetRepo) GetAll(ctx context.Context) ([]*models.Asset, error) {
	return models.Assets(qm.Where("date_deleted is null")).All(ctx, db.DB)
}

func (a AssetRepo) Insert(ctx context.Context, asset *models.Asset) error {
	return asset.Insert(ctx, db.DB, boil.Infer())
}
