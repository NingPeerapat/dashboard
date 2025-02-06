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

type ChartPtRepo struct {
	colName *mongo.Collection
	colTemp *mongo.Collection
}

func NewChartPtRepo(colName *mongo.Collection, colTemp *mongo.Collection) *ChartPtRepo {
	return &ChartPtRepo{
		colName: colName,
		colTemp: colTemp,
	}
}

func (repo *ChartPtRepo) GetChartUidData(body dto.ChartRequest) (*dao.UidData, error) {
	var data dao.UidData

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
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []dao.UidData
	for cursor.Next(ctx) {
		var result dao.UidData
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

func (repo *ChartPtRepo) GetChartPtTempData() ([]*dto.ChartPatientData, error) {
	var data []dao.ChartPatientTempData

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"chart_patient": 1,
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

	var result []*dto.ChartPatientData
	for _, d := range data {
		for _, patient := range d.ChartPatientTemp {
			result = append(result, &dto.ChartPatientData{
				DiseaseName:  patient.DiseaseName,
				QtyOfPatient: patient.QtyOfPatient,
				Avg:          patient.Avg,
			})
		}
	}

	return result, nil
}
