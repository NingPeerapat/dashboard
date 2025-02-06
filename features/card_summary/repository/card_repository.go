package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/card_summary/entities/dao"
	"ning/go-dashboard/features/card_summary/entities/dto"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardRepo struct {
	colName *mongo.Collection
	colTemp *mongo.Collection
}

func NewCardRepo(colName *mongo.Collection, colTemp *mongo.Collection) *CardRepo {
	return &CardRepo{
		colName: colName,
		colTemp: colTemp,
	}
}

func (repo *CardRepo) GetCardData(body dto.CardRequest) (*dao.CardRawData, error) {
	var data dao.CardRawData

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
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []dao.CardRawData
	for cursor.Next(ctx) {
		var result dao.CardRawData
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

func (repo *CardRepo) GetCidCountData(body dto.CardRequest) (*dao.CidCountData, error) {
	var data dao.CidCountData

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
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []dao.CidCountData
	for cursor.Next(ctx) {
		var result dao.CidCountData
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

func (repo *CardRepo) GetCradTempData() ([]*dto.CardData, error) {
	var data []dao.CardTempData

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"card_summary": 1,
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

	var result []*dto.CardData
	for _, d := range data {
		for _, card := range d.CardData {
			result = append(result, &dto.CardData{
				ServiceCount: card.ServiceCount,
				PatientCount: card.PatientCount,
				Expense:      card.Expense,
				HcodeCount:   card.HcodeCount,
				AvgService:   card.AvgService,
				AvgExpense:   card.AvgExpense,
			})
		}
	}

	return result, nil
}
