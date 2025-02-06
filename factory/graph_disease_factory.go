package factory

import (
	"ning/go-dashboard/features/graph_disease/controller"
	"ning/go-dashboard/features/graph_disease/repository"
	"ning/go-dashboard/features/graph_disease/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type GraphDiseaseCtrlFtr struct {
	GraphDiseasePtCtrl *controller.GraphDiseasePtCtrl
	GraphDiseaseExCtrl *controller.GraphDiseaseExCtrl
}

func NewGraphDiseaseCtrlFtr(client *mongo.Client, dbName, colName string) *GraphDiseaseCtrlFtr {
	graphDiseasePtRepo := repository.NewGraphDiseasePtRepo(client, dbName, colName)
	graphDiseasePtService := service.NewGraphDiseasePtService(graphDiseasePtRepo)
	graphDiseasePtCtrl := controller.NewGraphDiseasePtCtrl(graphDiseasePtService)

	graphDiseaseExRepo := repository.NewGraphDiseaseExRepo(client, dbName, colName)
	graphDiseaseExService := service.NewGraphDiseaseExService(graphDiseaseExRepo)
	graphDiseaseExCtrl := controller.NewGraphDiseaseExCtrl(graphDiseaseExService)

	return &GraphDiseaseCtrlFtr{
		GraphDiseasePtCtrl: graphDiseasePtCtrl,
		GraphDiseaseExCtrl: graphDiseaseExCtrl,
	}
}
