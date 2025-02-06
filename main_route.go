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

	graph_dm_c "ning/go-dashboard/features/graph_dm/controller"
	graph_dm_r "ning/go-dashboard/features/graph_dm/route"

	"github.com/gofiber/fiber/v2"
)

func RegisterAllRoutes(app *fiber.App,
	chartPtCtrl *chart_c.ChartPtCtrl,
	chartExCtrl *chart_c.ChartExCtrl,
	cardCtrl *card_c.CardCtrl,
	graphCaPtCtrl *graph_ca_c.GraphCaPtCtrl,
	graphCaExCtrl *graph_ca_c.GraphCaExCtrl,
	graphDiseasePtCtrl *graph_disease_c.GraphDiseasePtCtrl,
	graphDiseaseExCtrl *graph_disease_c.GraphDiseaseExCtrl,
	graphDmPtCtrl *graph_dm_c.GraphDmPtCtrl,
	graphDmExCtrl *graph_dm_c.GraphDmExCtrl) {

	chart_r.RegisterPatientRoutes(app, chartPtCtrl)
	chart_r.RegisterExpenseRoutes(app, chartExCtrl)
	card_r.RegisterCardRoutes(app, cardCtrl)
	graph_ca_r.RegisterPatientRoutes(app, graphCaPtCtrl)
	graph_ca_r.RegisterExpenseRoutes(app, graphCaExCtrl)
	graph_disease_r.RegisterPatientRoutes(app, graphDiseasePtCtrl)
	graph_disease_r.RegisterExpenseRoutes(app, graphDiseaseExCtrl)
	graph_dm_r.RegisterPatientRoutes(app, graphDmPtCtrl)
	graph_dm_r.RegisterExpenseRoutes(app, graphDmExCtrl)
}
