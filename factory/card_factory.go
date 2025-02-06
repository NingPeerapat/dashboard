package factory

import (
	"ning/go-dashboard/features/card_summary/controller"
	"ning/go-dashboard/features/card_summary/repository"
	"ning/go-dashboard/features/card_summary/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type CardCtrlFtr struct {
	CardCtrl *controller.CardCtrl
}

func NewCardCtrlFtr(client *mongo.Client, dbName, colName string) *CardCtrlFtr {
	cardRepo := repository.NewCardRepo(client, dbName, colName)
	cardService := service.NewCardService(cardRepo)
	cardCtrl := controller.NewCardCtrl(cardService)

	return &CardCtrlFtr{
		CardCtrl: cardCtrl,
	}
}
