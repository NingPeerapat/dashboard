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

type GraphDmPtRepo struct {
	colName *mongo.Collection
	colTemp *mongo.Collection
}

func NewGraphDmPtRepo(colName *mongo.Collection, colTemp *mongo.Collection) *GraphDmPtRepo {
	return &GraphDmPtRepo{
		colName: colName,
		colTemp: colTemp,
	}
}

func (repo *GraphDmPtRepo) GetDmPatient(body dto.DmRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":          0,
				"year":         1,
				"month":        1,
				"dm_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$dm_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":         "$year",
					"month":        "$month",
					"dm_cid_count": "$dm_cid_count",
				},
			},
		},
	}

	countStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":  "$_id.year",
					"month": "$_id.month",
				},
				"dm_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$dm_cid_count",
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

	pipeline := mongo.Pipeline{matchStage, projectStage1, unwindStage, groupStage, countStage, projectStage2, sortStage}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []utils.PatientData
	for cursor.Next(ctx) {
		var result utils.PatientData
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

func (repo *GraphDmPtRepo) GetHgPatient(body dto.DmRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":          0,
				"year":         1,
				"month":        1,
				"hg_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$hg_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":         "$year",
					"month":        "$month",
					"hg_cid_count": "$hg_cid_count",
				},
			},
		},
	}

	countStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":  "$_id.year",
					"month": "$_id.month",
				},
				"hg_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$hg_cid_count",
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

	pipeline := mongo.Pipeline{matchStage, projectStage1, unwindStage, groupStage, countStage, projectStage2, sortStage}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []utils.PatientData
	for cursor.Next(ctx) {
		var result utils.PatientData
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

func (repo *GraphDmPtRepo) GetDmCkdPatient(body dto.DmRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":              0,
				"year":             1,
				"month":            1,
				"dm_ckd_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$dm_ckd_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":             "$year",
					"month":            "$month",
					"dm_ckd_cid_count": "$dm_ckd_cid_count",
				},
			},
		},
	}

	countStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":  "$_id.year",
					"month": "$_id.month",
				},
				"dm_ckd_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$dm_ckd_cid_count",
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

	pipeline := mongo.Pipeline{matchStage, projectStage1, unwindStage, groupStage, countStage, projectStage2, sortStage}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []utils.PatientData
	for cursor.Next(ctx) {
		var result utils.PatientData
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

func (repo *GraphDmPtRepo) GetDmAcsPatient(body dto.DmRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":              0,
				"year":             1,
				"month":            1,
				"dm_acs_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$dm_acs_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":             "$year",
					"month":            "$month",
					"dm_acs_cid_count": "$dm_acs_cid_count",
				},
			},
		},
	}

	countStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":  "$_id.year",
					"month": "$_id.month",
				},
				"dm_acs_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$dm_acs_cid_count",
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

	pipeline := mongo.Pipeline{matchStage, projectStage1, unwindStage, groupStage, countStage, projectStage2, sortStage}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []utils.PatientData
	for cursor.Next(ctx) {
		var result utils.PatientData
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

func (repo *GraphDmPtRepo) GetDmCvaPatient(body dto.DmRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":              0,
				"year":             1,
				"month":            1,
				"dm_cva_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$dm_cva_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":             "$year",
					"month":            "$month",
					"dm_cva_cid_count": "$dm_cva_cid_count",
				},
			},
		},
	}

	countStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":  "$_id.year",
					"month": "$_id.month",
				},
				"dm_cva_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$dm_cva_cid_count",
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

	pipeline := mongo.Pipeline{matchStage, projectStage1, unwindStage, groupStage, countStage, projectStage2, sortStage}

	ctx := context.TODO()
	cursor, err := repo.colName.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var results []utils.PatientData
	for cursor.Next(ctx) {
		var result utils.PatientData
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

func (repo *GraphDmPtRepo) GetGraphDmPtTempData() ([]*dto.DmData, error) {
	var data []dao.GraphDmPtTempData

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"dm_patient": 1,
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

	var result []*dto.DmData
	for _, d := range data {
		for _, data := range d.GraphDmPtpenseTemp {
			result = append(result, &dto.DmData{
				DiseaseName: data.DiseaseName,
				Data:        data.Data,
			})
		}
	}

	return result, nil
}
