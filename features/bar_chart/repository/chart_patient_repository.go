package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/bar_chart/entities"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChartPatientRepository struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewChartPatientRepository(client *mongo.Client, databaseName string, collectionName string) *ChartPatientRepository {
	return &ChartPatientRepository{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (repo *ChartPatientRepository) GetAllData(body entities.ChartRequest) (*entities.UidData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)
	var data entities.UidData

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	groupStage := bson.D{
		{Key: "$group", Value: bson.M{
			"_id":              nil,
			"dm_uid_count":     bson.M{"$sum": "$dm_uid_count"},
			"ht_uid_count":     bson.M{"$sum": "$ht_uid_count"},
			"copd_uid_count":   bson.M{"$sum": "$copd_uid_count"},
			"ca_uid_count":     bson.M{"$sum": "$ca_uid_count"},
			"psy_uid_count":    bson.M{"$sum": "$psy_uid_count"},
			"hd_cvd_uid_count": bson.M{"$sum": "$hd_cvd_uid_count"},
		}},
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"_id":              0,
			"dm_uid_count":     1,
			"ht_uid_count":     1,
			"copd_uid_count":   1,
			"ca_uid_count":     1,
			"psy_uid_count":    1,
			"hd_cvd_uid_count": 1,
		}},
	}

	pipeline := mongo.Pipeline{matchStage, groupStage, projectStage}

	ctx := context.TODO()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []entities.UidData
	for cursor.Next(ctx) {
		var result entities.UidData
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
