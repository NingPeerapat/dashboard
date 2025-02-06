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

func NewGraphDiseaseCtrlFtr(colName *mongo.Collection, colTemp *mongo.Collection) *GraphDiseaseCtrlFtr {
	graphDiseasePtRepo := repository.NewGraphDiseasePtRepo(colName, colTemp)
	graphDiseasePtService := service.NewGraphDiseasePtService(graphDiseasePtRepo)
	graphDiseasePtCtrl := controller.NewGraphDiseasePtCtrl(graphDiseasePtService)

	graphDiseaseExRepo := repository.NewGraphDiseaseExRepo(colName, colTemp)
	graphDiseaseExService := service.NewGraphDiseaseExService(graphDiseaseExRepo)
	graphDiseaseExCtrl := controller.NewGraphDiseaseExCtrl(graphDiseaseExService)

	return &GraphDiseaseCtrlFtr{
		GraphDiseasePtCtrl: graphDiseasePtCtrl,
		GraphDiseaseExCtrl: graphDiseaseExCtrl,
	}
}
