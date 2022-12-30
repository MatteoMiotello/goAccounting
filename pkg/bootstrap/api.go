package bootstrap

import (
	"github.com/MatteoMiotello/goAccounting/pkg/api"
	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (a Api) Init() error {
	engine := gin.Default()
	baseRepo := api.Repository{Api: engine}

	err := api.AssetsRepository{Repository: baseRepo}.HandleRequests()
	if err != nil {
		return err
	}

	err = engine.Run()
	if err != nil {
		return err
	}

	return nil
}
