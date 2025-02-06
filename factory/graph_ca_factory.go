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

func NewGraphCaCtrlFtr(colName *mongo.Collection, colTemp *mongo.Collection) *GraphCaCtrlFtr {
	graphCaPtRepo := repository.NewGraphCaPtRepo(colName, colTemp)
	graphCaPtService := service.NewGraphCaPtService(graphCaPtRepo)
	graphCaPtCtrl := controller.NewGraphCaPtCtrl(graphCaPtService)

	graphCaExRepo := repository.NewGraphCaExRepo(colName, colTemp)
	graphCaExService := service.NewGraphCaExService(graphCaExRepo)
	graphCaExCtrl := controller.NewGraphCaExCtrl(graphCaExService)

	return &GraphCaCtrlFtr{
		GraphCaPtCtrl: graphCaPtCtrl,
		GraphCaExCtrl: graphCaExCtrl,
	}
}
