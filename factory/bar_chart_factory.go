package factory

import (
	"ning/go-dashboard/features/bar_chart/controller"
	"ning/go-dashboard/features/bar_chart/repository"
	"ning/go-dashboard/features/bar_chart/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type ChartCtrlFtr struct {
	ChartPtCtrl *controller.ChartPtCtrl
	ChartExCtrl *controller.ChartExCtrl
}

func NewChartCtrlFtr(colName *mongo.Collection, colTemp *mongo.Collection) *ChartCtrlFtr {
	// Chart Patient
	chartPtRepo := repository.NewChartPtRepo(colName, colTemp)
	chartPtCount := repository.NewCountCidRepo(colName, colTemp)
	chartPtService := service.NewChartPtService(chartPtRepo, chartPtCount)
	chartPtCtrl := controller.NewChartPtCtrl(chartPtService)

	// Chart Expense
	chartExRepo := repository.NewChartExRepo(colName, colTemp)
	chartExCount := repository.NewCountCidRepo(colName, colTemp)
	chartExService := service.NewChartExService(chartExRepo, chartExCount)
	chartExCtrl := controller.NewChartExCtrl(chartExService)

	return &ChartCtrlFtr{
		ChartPtCtrl: chartPtCtrl,
		ChartExCtrl: chartExCtrl,
	}
}
