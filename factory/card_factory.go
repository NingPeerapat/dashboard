package factory

import (
	"ning/go-dashboard/features/card_summary/controller"
	"ning/go-dashboard/features/card_summary/repository"
	"ning/go-dashboard/features/card_summary/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type CardControllerFactory struct {
	CardController *controller.CardController
}

func NewCardControllerFactory(client *mongo.Client, dbName, colName string) *CardControllerFactory {
	repoCard := repository.NewCardRepository(client, dbName, colName)
	cardService := service.NewCardService(repoCard)
	cardController := controller.NewCardController(cardService)

	return &CardControllerFactory{
		CardController: cardController,
	}
}
