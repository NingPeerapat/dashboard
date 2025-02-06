package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/graph_dm/entities/dao"
	"ning/go-dashboard/features/graph_dm/entities/dto"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GraphDmExRepo struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewGraphDmExRepo(client *mongo.Client, databaseName string, collectionName string) *GraphDmExRepo {
	return &GraphDmExRepo{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (repo *GraphDmExRepo) GetGraphDmExData(body dto.DmRequest) ([]dao.DmExpenseData, error) {
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
				"hg_expense":     bson.M{"$sum": "$hg_expense"},
				"dm_ckd_expense": bson.M{"$sum": "$dm_ckd_expense"},
				"dm_acs_expense": bson.M{"$sum": "$dm_acs_expense"},
				"dm_cva_expense": bson.M{"$sum": "$dm_cva_expense"},
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
				"hg_expense":     1,
				"dm_ckd_expense": 1,
				"dm_acs_expense": 1,
				"dm_cva_expense": 1,
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

	var results []dao.DmExpenseData
	for cursor.Next(ctx) {
		var result dao.DmExpenseData
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
