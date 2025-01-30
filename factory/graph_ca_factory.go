package factory

import (
	"ning/go-dashboard/features/graph_ca/controller"
	"ning/go-dashboard/features/graph_ca/repository"
	"ning/go-dashboard/features/graph_ca/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type GraphCaControllerFactory struct {
	GraphCaPatientController *controller.GraphCaPatientController
	GraphCaExpenseController *controller.GraphCaExpenseController
}

func NewGraphCaControllerFactory(client *mongo.Client, dbName, colName string) *GraphCaControllerFactory {
	repoGraphCaPatient := repository.NewGraphCaPatientRepository(client, dbName, colName)
	graphCaPatientService := service.NewGraphCaPatientService(repoGraphCaPatient)
	graphCaPatientController := controller.NewGraphCaPatientController(graphCaPatientService)

	repoGraphCaExpense := repository.NewGraphCaExpenseRepository(client, dbName, colName)
	graphCaExpenseService := service.NewGraphCaExpenseService(repoGraphCaExpense)
	graphCaExpenseController := controller.NewGraphCaExpenseController(graphCaExpenseService)

	return &GraphCaControllerFactory{
		GraphCaPatientController: graphCaPatientController,
		GraphCaExpenseController: graphCaExpenseController,
	}
}
