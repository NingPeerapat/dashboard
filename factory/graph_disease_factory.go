package factory

import (
	"ning/go-dashboard/features/graph_disease/controller"
	"ning/go-dashboard/features/graph_disease/repository"
	"ning/go-dashboard/features/graph_disease/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type DiseaseControllerFactory struct {
	DiseasePatientController *controller.DiseasePatientController
	DiseaseExpenseController *controller.DiseaseExpenseController
}

func NewDiseaseControllerFactory(client *mongo.Client, dbName, colName string) *DiseaseControllerFactory {
	repoDiseasePatient := repository.NewDiseasePatientRepository(client, dbName, colName)
	diseasePatientService := service.NewDiseasePatientService(repoDiseasePatient)
	diseasePatientController := controller.NewDiseasePatientController(diseasePatientService)

	repoDiseaseExpense := repository.NewDiseaseExpenseRepository(client, dbName, colName)
	diseaseExpenseService := service.NewDiseaseExpenseService(repoDiseaseExpense)
	diseaseExpenseController := controller.NewDiseaseExpenseController(diseaseExpenseService)

	return &DiseaseControllerFactory{
		DiseasePatientController: diseasePatientController,
		DiseaseExpenseController: diseaseExpenseController,
	}
}
