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

type GraphCaPtRepo struct {
	colName *mongo.Collection
	colTemp *mongo.Collection
}

func NewGraphCaPtRepo(colName *mongo.Collection, colTemp *mongo.Collection) *GraphCaPtRepo {
	return &GraphCaPtRepo{
		colName: colName,
		colTemp: colTemp,
	}
}

func (repo *GraphCaPtRepo) GetCaPatient(body dto.CaRequest) ([]utils.PatientData, error) {

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
				"ca_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$ca_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":         "$year",
					"month":        "$month",
					"ca_cid_count": "$ca_cid_count",
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
				"ca_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$ca_cid_count",
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

func (repo *GraphCaPtRepo) GetLungCaPatient(body dto.CaRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":               0,
				"year":              1,
				"month":             1,
				"lung_ca_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$lung_ca_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":              "$year",
					"month":             "$month",
					"lung_ca_cid_count": "$lung_ca_cid_count",
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
				"lung_ca_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$lung_ca_cid_count",
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

func (repo *GraphCaPtRepo) GetBreastCaPatient(body dto.CaRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":                 0,
				"year":                1,
				"month":               1,
				"breast_ca_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$breast_ca_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":                "$year",
					"month":               "$month",
					"breast_ca_cid_count": "$breast_ca_cid_count",
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
				"breast_ca_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$breast_ca_cid_count",
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

func (repo *GraphCaPtRepo) GetCervicalCaPatient(body dto.CaRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":                   0,
				"year":                  1,
				"month":                 1,
				"cervical_ca_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$cervical_ca_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":                  "$year",
					"month":                 "$month",
					"cervical_ca_cid_count": "$cervical_ca_cid_count",
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
				"cervical_ca_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$cervical_ca_cid_count",
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

func (repo *GraphCaPtRepo) GetLiverCaPatient(body dto.CaRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":                0,
				"year":               1,
				"month":              1,
				"liver_ca_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$liver_ca_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":               "$year",
					"month":              "$month",
					"liver_ca_cid_count": "$liver_ca_cid_count",
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
				"liver_ca_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$liver_ca_cid_count",
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

func (repo *GraphCaPtRepo) GetColorectalCaPatient(body dto.CaRequest) ([]utils.PatientData, error) {

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":                     0,
				"year":                    1,
				"month":                   1,
				"colorectal_ca_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$colorectal_ca_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":                    "$year",
					"month":                   "$month",
					"colorectal_ca_cid_count": "$colorectal_ca_cid_count",
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
				"colorectal_ca_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$colorectal_ca_cid_count",
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

func (repo *GraphCaPtRepo) GetGraphCaPtTempData() ([]*dto.CaData, error) {
	var data []dao.GraphCaPtTempData

	pipeline := []bson.M{
		{
			"$project": bson.M{
				"ca_patient": 1,
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
		for _, data := range d.GraphCaPtpenseTemp {
			result = append(result, &dto.CaData{
				DiseaseName: data.DiseaseName,
				Data:        data.Data,
			})
		}
	}

	return result, nil
}
