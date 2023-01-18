package main

import (
	"github.com/MatteoMiotello/goAccounting/internal/api"
	"github.com/MatteoMiotello/goAccounting/internal/bootstrap"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitDb()
	bootstrap.InitModels()
}

func main() {
	api.Run()
}
