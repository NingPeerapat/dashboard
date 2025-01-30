package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/bar_chart/entities"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChartExpenseRepository struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewChartExpenseRepository(client *mongo.Client, databaseName string, collectionName string) *ChartExpenseRepository {
	return &ChartExpenseRepository{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (repo *ChartExpenseRepository) GetAllData(body entities.ChartRequest) (*entities.ExpenseData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)
	var data entities.ExpenseData

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	groupStage := bson.D{
		{Key: "$group", Value: bson.M{
			"_id":            nil,
			"dm_expense":     bson.M{"$sum": "$dm_expense"},
			"ht_expense":     bson.M{"$sum": "$ht_expense"},
			"copd_expense":   bson.M{"$sum": "$copd_expense"},
			"ca_expense":     bson.M{"$sum": "$ca_expense"},
			"psy_expense":    bson.M{"$sum": "$psy_expense"},
			"hd_cvd_expense": bson.M{"$sum": "$hd_cvd_expense"},
		}},
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"_id":            0,
			"dm_expense":     1,
			"ht_expense":     1,
			"copd_expense":   1,
			"ca_expense":     1,
			"psy_expense":    1,
			"hd_cvd_expense": 1,
		}},
	}

	pipeline := mongo.Pipeline{matchStage, groupStage, projectStage}

	ctx := context.TODO()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []entities.ExpenseData
	for cursor.Next(ctx) {
		var result entities.ExpenseData
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding data: %v", err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	if len(results) > 0 {
		data = results[0]
	}

	return &data, nil
}
