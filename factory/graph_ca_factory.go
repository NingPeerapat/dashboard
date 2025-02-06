package factory

import (
	"ning/go-dashboard/features/graph_ca/controller"
	"ning/go-dashboard/features/graph_ca/repository"
	"ning/go-dashboard/features/graph_ca/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type GraphCaCtrlFtr struct {
	GraphCaPtCtrl *controller.GraphCaPtCtrl
	GraphCaExCtrl *controller.GraphCaExCtrl
}

func NewGraphCaCtrlFtr(client *mongo.Client, dbName, colName string) *GraphCaCtrlFtr {
	graphCaPtRepo := repository.NewGraphCaPtRepo(client, dbName, colName)
	graphCaPtService := service.NewGraphCaPtService(graphCaPtRepo)
	graphCaPtCtrl := controller.NewGraphCaPtCtrl(graphCaPtService)

	graphCaExRepo := repository.NewGraphCaExRepo(client, dbName, colName)
	graphCaExService := service.NewGraphCaExService(graphCaExRepo)
	graphCaExCtrl := controller.NewGraphCaExCtrl(graphCaExService)

	return &GraphCaCtrlFtr{
		GraphCaPtCtrl: graphCaPtCtrl,
		GraphCaExCtrl: graphCaExCtrl,
	}
}
