package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/bar_chart/entities/dao"
	"ning/go-dashboard/features/bar_chart/entities/dto"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChartExRepo struct {
	colName *mongo.Collection
	colTemp *mongo.Collection
}

func NewChartExRepo(colName *mongo.Collection, colTemp *mongo.Collection) *ChartExRepo {
	return &ChartExRepo{
		colName: colName,
		colTemp: colTemp,
	}
}

func (repo *ChartExRepo) GetChartExData(body dto.ChartRequest) (*dao.ExpenseData, error) {
	var data dao.ExpenseData

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
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []dao.ExpenseData
	for cursor.Next(ctx) {
		var result dao.ExpenseData
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

func (repo *ChartExRepo) GetChartExTempData() ([]*dto.ChartExpenseData, error) {
	var data []dao.ChartExpenseTempData

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"chart_expense": 1,
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

	var result []*dto.ChartExpenseData
	for _, d := range data {
		for _, expense := range d.ChartExpenseTemp {
			result = append(result, &dto.ChartExpenseData{
				DiseaseName:  expense.DiseaseName,
				QtyOfExpense: expense.QtyOfExpense,
				Avg:          expense.Avg,
			})
		}
	}

	return result, nil
}
