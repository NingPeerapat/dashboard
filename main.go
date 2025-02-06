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
	client, err := database.ConnectMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	chartFtr := factory.NewChartCtrlFtr(client, cfg.DatabaseName, cfg.CollectionName)
	cardFtr := factory.NewCardCtrlFtr(client, cfg.DatabaseName, cfg.CollectionName)
	graphCaFtr := factory.NewGraphCaCtrlFtr(client, cfg.DatabaseName, cfg.CollectionName)
	graphDiseaseFtr := factory.NewGraphDiseaseCtrlFtr(client, cfg.DatabaseName, cfg.CollectionName)
	graphDmFtr := factory.NewGraphDmCtrlFtr(client, cfg.DatabaseName, cfg.CollectionName)

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
