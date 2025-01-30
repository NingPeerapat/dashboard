package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/card_summary/entities"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardRepository struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewCardRepository(client *mongo.Client, databaseName string, collectionName string) *CardRepository {
	return &CardRepository{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (repo *CardRepository) GetAllData(body entities.CardRequest) (*entities.CardRawData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)
	var data entities.CardRawData

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	groupStage := bson.D{
		{Key: "$group", Value: bson.M{
			"_id":         nil,
			"hcode_count": bson.M{"$addToSet": "$hcode"},
			"service_count": bson.M{
				"$sum": bson.M{
					"$add": []string{
						"$dm_uid_count",
						"$ht_uid_count",
						"$copd_uid_count",
						"$ca_uid_count",
						"$psy_uid_count",
						"$hd_cvd_uid_count",
					},
				},
			},
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
			"hcode_count": bson.M{
				"$size": "$hcode_count",
			},
			"service_count":  1,
			"dm_expense":     1,
			"ht_expense":     1,
			"copd_expense":   1,
			"ca_expense":     1,
			"psy_expense":    1,
			"hd_cvd_expense": 1,
		},
		},
	}

	pipeline := mongo.Pipeline{
		matchStage,
		groupStage,
		projectStage,
	}

	ctx := context.TODO()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []entities.CardRawData
	for cursor.Next(ctx) {
		var result entities.CardRawData
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

func (repo *CardRepository) GetCidCountData(body entities.CardRequest) (*entities.CidCountData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)
	var data entities.CidCountData

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"cid_list": bson.M{
				"$setUnion": []string{"$dm_cid_count", "$ht_cid_count", "$copd_cid_count", "$ca_cid_count", "$psy_cid_count", "$hd_cvd_cid_count"},
			},
		},
		},
	}

	pipeline := mongo.Pipeline{
		matchStage,
		projectStage,
		bson.D{
			{Key: "$unwind", Value: "$cid_list"},
		},
		bson.D{
			{Key: "$group", Value: bson.M{"_id": "$cid_list"}},
		},
		bson.D{
			{Key: "$count", Value: "cid_count"},
		},
	}

	ctx := context.TODO()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []entities.CidCountData
	for cursor.Next(ctx) {
		var result entities.CidCountData
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
