package utils

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func MatchStageCardBar(startDate time.Time, endDate time.Time, area string, province string, district string, hcode string) (bson.D, error) {

	matchFilter := bson.M{
		"date": bson.M{
			"$gte": startDate,
			"$lte": endDate,
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
