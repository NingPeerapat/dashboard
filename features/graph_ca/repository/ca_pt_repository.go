package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/graph_ca/entities/dto"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GraphCaPtRepo struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewGraphCaPtRepo(client *mongo.Client, databaseName string, collectionName string) *GraphCaPtRepo {
	return &GraphCaPtRepo{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (repo *GraphCaPtRepo) GetCaPatient(body dto.CaRequest) ([]utils.PatientData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

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
	cursor, err := collection.Aggregate(ctx, pipeline)
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
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

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
	cursor, err := collection.Aggregate(ctx, pipeline)
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
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

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
	cursor, err := collection.Aggregate(ctx, pipeline)
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
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

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
	cursor, err := collection.Aggregate(ctx, pipeline)
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
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

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
	cursor, err := collection.Aggregate(ctx, pipeline)
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
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

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
	cursor, err := collection.Aggregate(ctx, pipeline)
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
