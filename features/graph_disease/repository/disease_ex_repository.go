package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/graph_disease/entities/dao"
	"ning/go-dashboard/features/graph_disease/entities/dto"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GraphDiseaseExRepo struct {
	colName *mongo.Collection
	colTemp *mongo.Collection
}

func NewGraphDiseaseExRepo(colName *mongo.Collection, colTemp *mongo.Collection) *GraphDiseaseExRepo {
	return &GraphDiseaseExRepo{
		colName: colName,
		colTemp: colTemp,
	}
}

func (repo *GraphDiseaseExRepo) GetGraphDiseaseExData(body dto.DiseaseRequest) ([]dao.DiseaseExpenseData, error) {

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
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []dao.DiseaseExpenseData
	for cursor.Next(ctx) {
		var result dao.DiseaseExpenseData
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

func (repo *GraphDiseaseExRepo) GetGraphDiseaseExTempData() ([]*dto.DiseaseData, error) {
	var data []dao.GraphDiseaseExTempData

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"disease_expense": 1,
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

	var result []*dto.DiseaseData
	for _, d := range data {
		for _, data := range d.GraphDiseaseExpenseTemp {
			result = append(result, &dto.DiseaseData{
				DiseaseName: data.DiseaseName,
				Data:        data.Data,
			})
		}
	}

	return result, nil
}
