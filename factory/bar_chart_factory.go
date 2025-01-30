package factory

import (
	"ning/go-dashboard/features/bar_chart/controller"
	"ning/go-dashboard/features/bar_chart/repository"
	"ning/go-dashboard/features/bar_chart/service"
	"ning/go-dashboard/pkg/shared"

	"go.mongodb.org/mongo-driver/mongo"
)

type ChartControllerFactory struct {
	ChartExpenseController *controller.ChartExpenseController
	ChartPatientController *controller.ChartPatientController
}

func NewChartControllerFactory(client *mongo.Client, dbName, colName string) *ChartControllerFactory {
	// Chart Expense
	repoChartEx := repository.NewChartExpenseRepository(client, dbName, colName)
	countChartEx := shared.NewCountCidRepository(client, dbName, colName)
	chartExpenseService := service.NewChartExpenseService(repoChartEx, countChartEx)
	chartExpenseController := controller.NewChartExpenseController(chartExpenseService)

	// Chart Patient
	repoChartPt := repository.NewChartPatientRepository(client, dbName, colName)
	countChartPt := shared.NewCountCidRepository(client, dbName, colName)
	chartPatientService := service.NewChartPatientService(repoChartPt, countChartPt)
	chartPatientController := controller.NewChartPatientController(chartPatientService)

	return &ChartControllerFactory{
		ChartExpenseController: chartExpenseController,
		ChartPatientController: chartPatientController,
	}
}
