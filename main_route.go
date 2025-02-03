package main

import (
	chart_c "ning/go-dashboard/features/bar_chart/controller"
	chart_r "ning/go-dashboard/features/bar_chart/route"

	card_c "ning/go-dashboard/features/card_summary/controller"
	card_r "ning/go-dashboard/features/card_summary/route"

	graph_ca_c "ning/go-dashboard/features/graph_ca/controller"
	graph_ca_r "ning/go-dashboard/features/graph_ca/route"

	graph_disease_c "ning/go-dashboard/features/graph_disease/controller"
	graph_disease_r "ning/go-dashboard/features/graph_disease/route"

	"github.com/gofiber/fiber/v2"
)

func RegisterAllRoutes(app *fiber.App,
	chartExController *chart_c.ChartExpenseController,
	chartPtController *chart_c.ChartPatientController,
	cardController *card_c.CardController,
	graphCaPatientController *graph_ca_c.GraphCaPatientController,
	graphCaExpenseController *graph_ca_c.GraphCaExpenseController,
	diseasePatientController *graph_disease_c.DiseasePatientController,
	diseaseExpenseController *graph_disease_c.DiseaseExpenseController) {

	chart_r.RegisterExpenseRoutes(app, chartExController)
	chart_r.RegisterPatientRoutes(app, chartPtController)
	card_r.RegisterCardRoutes(app, cardController)
	graph_ca_r.RegisterPatientRoutes(app, graphCaPatientController)
	graph_ca_r.RegisterExpenseRoutes(app, graphCaExpenseController)
	graph_disease_r.RegisterPatientRoutes(app, diseasePatientController)
	graph_disease_r.RegisterExpenseRoutes(app, diseaseExpenseController)
}
