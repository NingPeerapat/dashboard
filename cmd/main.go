package main

import (
	"fmt"
	"log"
	"ning/go-dashboard/factory"
	"ning/go-dashboard/pkg/database"
	l "ning/go-dashboard/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// โหลด config จาก .env
	cfg, err := database.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// เชื่อมต่อกับ MongoDB
	colName, colTempName, err := database.ConnectMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	chartFtr := factory.NewChartCtrlFtr(colName, colTempName)
	cardFtr := factory.NewCardCtrlFtr(colName, colTempName)
	graphCaFtr := factory.NewGraphCaCtrlFtr(colName, colTempName)
	graphDiseaseFtr := factory.NewGraphDiseaseCtrlFtr(colName, colTempName)
	graphDmFtr := factory.NewGraphDmCtrlFtr(colName, colTempName)

	app := fiber.New()

	app.Use(l.ApiLog())

	RegisterAllRoutes(app,
		chartFtr.ChartPtCtrl,
		chartFtr.ChartExCtrl,
		cardFtr.CardCtrl,
		graphCaFtr.GraphCaPtCtrl,
		graphCaFtr.GraphCaExCtrl,
		graphDiseaseFtr.GraphDiseasePtCtrl,
		graphDiseaseFtr.GraphDiseaseExCtrl,
		graphDmFtr.GraphDmPtCtrl,
		graphDmFtr.GraphDmExCtrl)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
