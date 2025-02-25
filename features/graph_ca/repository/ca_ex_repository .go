package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/graph_ca/entities/dao"
	"ning/go-dashboard/features/graph_ca/entities/dto"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GraphCaExRepo struct {
	colName *mongo.Collection
	colTemp *mongo.Collection
}

func NewGraphCaExRepo(colName *mongo.Collection, colTemp *mongo.Collection) *GraphCaExRepo {
	return &GraphCaExRepo{
		colName: colName,
		colTemp: colTemp,
	}
}

func (repo *GraphCaExRepo) GetGraphCaExData(body dto.CaRequest) ([]dao.CaExpenseData, error) {

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
				"ca_expense":            bson.M{"$sum": "$ca_expense"},
				"lung_ca_expense":       bson.M{"$sum": "$lung_ca_expense"},
				"breast_ca_expense":     bson.M{"$sum": "$breast_ca_expense"},
				"cervical_ca_expense":   bson.M{"$sum": "$cervical_ca_expense"},
				"liver_ca_expense":      bson.M{"$sum": "$liver_ca_expense"},
				"colorectal_ca_expense": bson.M{"$sum": "$colorectal_ca_expense"},
			},
		},
	}

	projectStage := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":                   0,
				"year":                  "$_id.year",
				"month":                 "$_id.month",
				"ca_expense":            1,
				"lung_ca_expense":       1,
				"breast_ca_expense":     1,
				"cervical_ca_expense":   1,
				"liver_ca_expense":      1,
				"colorectal_ca_expense": 1,
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
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []dao.CaExpenseData
	for cursor.Next(ctx) {
		var result dao.CaExpenseData
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

func (repo *GraphCaExRepo) GetGraphCaExTempData() ([]*dto.CaData, error) {
	var data []dao.GraphCaExTempData

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"ca_expense": 1,
			},
		},
	}

	ctx := context.TODO()

	cursor, err := repo.colTemp.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &data); err != nil {
		return nil, fmt.Errorf("error decoding data: %v", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	var result []*dto.CaData
	for _, d := range data {
		for _, data := range d.GraphCaExpenseTemp {
			result = append(result, &dto.CaData{
				DiseaseName: data.DiseaseName,
				Data:        data.Data,
			})
		}
	}

	return result, nil
}
