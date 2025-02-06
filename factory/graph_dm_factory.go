package factory

import (
	"ning/go-dashboard/features/graph_dm/controller"
	"ning/go-dashboard/features/graph_dm/repository"
	"ning/go-dashboard/features/graph_dm/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type GraphDmCtrlFtr struct {
	GraphDmPtCtrl *controller.GraphDmPtCtrl
	GraphDmExCtrl *controller.GraphDmExCtrl
}

func NewGraphDmCtrlFtr(colName *mongo.Collection, colTemp *mongo.Collection) *GraphDmCtrlFtr {
	graphDmPtRepo := repository.NewGraphDmPtRepo(colName, colTemp)
	graphDmPtService := service.NewGraphDmPtService(graphDmPtRepo)
	graphDmPtCtrl := controller.NewGraphDmPtCtrl(graphDmPtService)

	graphDmExRepo := repository.NewGraphDmExRepo(colName, colTemp)
	graphDmExService := service.NewGraphDmExService(graphDmExRepo)
	graphDmExCtrl := controller.NewGraphDmExCtrl(graphDmExService)

	return &GraphDmCtrlFtr{
		GraphDmPtCtrl: graphDmPtCtrl,
		GraphDmExCtrl: graphDmExCtrl,
	}
}
