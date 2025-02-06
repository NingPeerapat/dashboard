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

func NewGraphDmCtrlFtr(client *mongo.Client, dbName, colName string) *GraphDmCtrlFtr {
	graphDmPtRepo := repository.NewGraphDmPtRepo(client, dbName, colName)
	graphDmPtService := service.NewGraphDmPtService(graphDmPtRepo)
	graphDmPtCtrl := controller.NewGraphDmPtCtrl(graphDmPtService)

	graphDmExRepo := repository.NewGraphDmExRepo(client, dbName, colName)
	graphDmExService := service.NewGraphDmExService(graphDmExRepo)
	graphDmExCtrl := controller.NewGraphDmExCtrl(graphDmExService)

	return &GraphDmCtrlFtr{
		GraphDmPtCtrl: graphDmPtCtrl,
		GraphDmExCtrl: graphDmExCtrl,
	}
}
