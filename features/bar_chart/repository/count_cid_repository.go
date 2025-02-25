package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/bar_chart/entities/dto"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CountCidRepo struct {
	colName *mongo.Collection
	colTemp *mongo.Collection
}

func NewCountCidRepo(colName *mongo.Collection, colTemp *mongo.Collection) *CountCidRepo {
	return &CountCidRepo{
		colName: colName,
		colTemp: colTemp,
	}
}

func (repo *CountCidRepo) CountDmCid(body dto.ChartRequest) (int, error) {

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return 0, err
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"dm_cid_count": bson.M{
				"$setUnion": bson.A{"$dm_cid_count"},
			},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, projectStage,
		{
			{Key: "$unwind", Value: "$dm_cid_count"},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$dm_cid_count",
			}},
		},
		{
			{Key: "$count", Value: "dm_cid_count"},
		}}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		DmCidCount int `bson:"dm_cid_count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("error decoding data: %v", err)
		}
	}

	return result.DmCidCount, nil
}

func (repo *CountCidRepo) CountHtCid(body dto.ChartRequest) (int, error) {

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return 0, err
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"ht_cid_count": bson.M{
				"$setUnion": bson.A{"$ht_cid_count"},
			},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, projectStage,
		{
			{Key: "$unwind", Value: "$ht_cid_count"},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$ht_cid_count",
			}},
		},
		{
			{Key: "$count", Value: "ht_cid_count"},
		}}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		HtCidCount int `bson:"ht_cid_count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("error decoding data: %v", err)
		}
	}

	return result.HtCidCount, nil
}

func (repo *CountCidRepo) CountCopdCid(body dto.ChartRequest) (int, error) {

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return 0, err
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"copd_cid_count": bson.M{
				"$setUnion": bson.A{"$copd_cid_count"},
			},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, projectStage,
		{
			{Key: "$unwind", Value: "$copd_cid_count"},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$copd_cid_count",
			}},
		},
		{
			{Key: "$count", Value: "copd_cid_count"},
		}}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		CopdCidCount int `bson:"copd_cid_count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("error decoding data: %v", err)
		}
	}

	return result.CopdCidCount, nil
}

func (repo *CountCidRepo) CountCaCid(body dto.ChartRequest) (int, error) {

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return 0, err
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"ca_cid_count": bson.M{
				"$setUnion": bson.A{"$ca_cid_count"},
			},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, projectStage,
		{
			{Key: "$unwind", Value: "$ca_cid_count"},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$ca_cid_count",
			}},
		},
		{
			{Key: "$count", Value: "ca_cid_count"},
		}}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		CaCidCount int `bson:"ca_cid_count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("error decoding data: %v", err)
		}
	}

	return result.CaCidCount, nil
}

func (repo *CountCidRepo) CountPsyCid(body dto.ChartRequest) (int, error) {

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return 0, err
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"psy_cid_count": bson.M{
				"$setUnion": bson.A{"$psy_cid_count"},
			},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, projectStage,
		{
			{Key: "$unwind", Value: "$psy_cid_count"},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$psy_cid_count",
			}},
		},
		{
			{Key: "$count", Value: "psy_cid_count"},
		}}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		PsyCidCount int `bson:"psy_cid_count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("error decoding data: %v", err)
		}
	}

	return result.PsyCidCount, nil
}

func (repo *CountCidRepo) CountHdCvdCid(body dto.ChartRequest) (int, error) {

	matchStage, err := utils.MatchStageCardBar(body.StartDate, body.EndDate, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return 0, err
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"hd_cvd_cid_count": bson.M{
				"$setUnion": bson.A{"$hd_cvd_cid_count"},
			},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, projectStage,
		{
			{Key: "$unwind", Value: "$hd_cvd_cid_count"},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$hd_cvd_cid_count",
			}},
		},
		{
			{Key: "$count", Value: "hd_cvd_cid_count"},
		}}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		HdCvdCidCount int `bson:"hd_cvd_cid_count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("error decoding data: %v", err)
		}
	}

	return result.HdCvdCidCount, nil
}
