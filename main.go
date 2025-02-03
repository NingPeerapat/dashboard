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

	chartFactory := factory.NewChartControllerFactory(client, cfg.DatabaseName, cfg.CollectionName)
	cardFactory := factory.NewCardControllerFactory(client, cfg.DatabaseName, cfg.CollectionName)
	graphCaFactory := factory.NewGraphCaControllerFactory(client, cfg.DatabaseName, cfg.CollectionName)
	graphDiseaseFactory := factory.NewDiseaseControllerFactory(client, cfg.DatabaseName, cfg.CollectionName)

	app := fiber.New()

	app.Use(l.ApiLog())

	RegisterAllRoutes(app,
		chartFactory.ChartExpenseController,
		chartFactory.ChartPatientController,
		cardFactory.CardController,
		graphCaFactory.GraphCaPatientController,
		graphCaFactory.GraphCaExpenseController,
		graphDiseaseFactory.DiseasePatientController,
		graphDiseaseFactory.DiseaseExpenseController)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
