package repository

import (
	"context"
	"fmt"
	"ning/go-dashboard/features/graph_disease/entities"
	"ning/go-dashboard/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DiseasePatientRepository struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

func NewDiseasePatientRepository(client *mongo.Client, databaseName string, collectionName string) *DiseasePatientRepository {
	return &DiseasePatientRepository{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

func (repo *DiseasePatientRepository) GetDmPatient(body entities.DiseaseRequest) ([]utils.PatientData, error) {
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
				"patient": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": 1,
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

func (repo *DiseasePatientRepository) GetHtPatient(body entities.DiseaseRequest) ([]utils.PatientData, error) {
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
				"ht_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$ht_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":         "$year",
					"month":        "$month",
					"ht_cid_count": "$ht_cid_count",
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
				"ht_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$ht_cid_count",
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

func (repo *DiseasePatientRepository) GetCopdPatient(body entities.DiseaseRequest) ([]utils.PatientData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":            0,
				"year":           1,
				"month":          1,
				"copd_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$copd_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":           "$year",
					"month":          "$month",
					"copd_cid_count": "$copd_cid_count",
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
				"copd_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$copd_cid_count",
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

func (repo *DiseasePatientRepository) GetCaPatient(body entities.DiseaseRequest) ([]utils.PatientData, error) {
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

func (repo *DiseasePatientRepository) GetPsyPatient(body entities.DiseaseRequest) ([]utils.PatientData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

	matchStage, err := utils.MatchStageGraph(body.Year, body.Area, body.Province, body.District, body.Hcode)
	if err != nil {
		return nil, err
	}

	projectStage1 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":           0,
				"year":          1,
				"month":         1,
				"psy_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$psy_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":          "$year",
					"month":         "$month",
					"psy_cid_count": "$psy_cid_count",
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
				"psy_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$psy_cid_count",
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

func (repo *DiseasePatientRepository) GetHdCvdPatient(body entities.DiseaseRequest) ([]utils.PatientData, error) {
	collection := repo.client.Database(repo.databaseName).Collection(repo.collectionName)

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
				"hd_cvd_cid_count": 1,
			},
		},
	}

	unwindStage := bson.D{{Key: "$unwind", Value: "$hd_cvd_cid_count"}}

	groupStage := bson.D{
		{Key: "$group",
			Value: bson.M{
				"_id": bson.M{
					"year":             "$year",
					"month":            "$month",
					"hd_cvd_cid_count": "$hd_cvd_cid_count",
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
				"hd_cvd_cid_count": bson.M{"$sum": 1},
			},
		},
	}

	projectStage2 := bson.D{
		{Key: "$project",
			Value: bson.M{
				"_id":     0,
				"year":    "$_id.year",
				"month":   "$_id.month",
				"patient": "$hd_cvd_cid_count",
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
