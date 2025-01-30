package utils

import (
	"go.mongodb.org/mongo-driver/bson"
)

func MatchStageGraph(year int, area string, province string, district string, hcode string) (bson.D, error) {

	matchFilter := bson.M{
		"$or": []bson.M{
			{"year": year, "month": bson.M{"$gte": 4}},
			{"year": year + 1, "month": bson.M{"$lte": 3}},
		},
	}

	if area != "" {
		matchFilter["area"] = area
	}
	if province != "" {
		matchFilter["province"] = province
	}
	if district != "" {
		matchFilter["district"] = district
	}
	if hcode != "" {
		matchFilter["hcode"] = hcode
	}

	matchStage := bson.D{{Key: "$match", Value: matchFilter}}

	return matchStage, nil
}
