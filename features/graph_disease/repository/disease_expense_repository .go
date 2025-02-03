package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/graph_disease/entities"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DiseaseExpenseRepository struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewDiseaseExpenseRepository(client *mongo.Client, databaseName string, collectionName string) *DiseaseExpenseRepository {
	return &DiseaseExpenseRepository{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (repo *DiseaseExpenseRepository) GetDiseaseExpense(body entities.DiseaseRequest) ([]entities.DiseaseExpenseData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":  "$year",
					"month": "$month",
				},
				"dm_expense":     bson.M{"$sum": "$dm_expense"},
				"ht_expense":     bson.M{"$sum": "$ht_expense"},
				"copd_expense":   bson.M{"$sum": "$copd_expense"},
				"ca_expense":     bson.M{"$sum": "$ca_expense"},
				"psy_expense":    bson.M{"$sum": "$psy_expense"},
				"hd_cvd_expense": bson.M{"$sum": "$hd_cvd_expense"},
			},
		},
	}

	projectStage := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":            0,
				"year":           "$_id.year",
				"month":          "$_id.month",
				"dm_expense":     1,
				"ht_expense":     1,
				"copd_expense":   1,
				"ca_expense":     1,
				"psy_expense":    1,
				"hd_cvd_expense": 1,
			},
		},
	}

	sortStage := bson.D{
		{Key: "$sort",
			Value: bson.M{
				"year":  1,
				"month": 1,
			},
		},
	}

	pipeline := mongo.Pipeline{matchStage, groupStage, projectStage, sortStage}

	ctx := context.TODO()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []entities.DiseaseExpenseData
	for cursor.Next(ctx) {
		var result entities.DiseaseExpenseData
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding data: %v", err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return results, nil
}
