package dao

type UidData struct {
	DmUidCount    float64 `bson:"dm_uid_count" json:"dm_uid_count"`
	HtUidCount    float64 `bson:"ht_uid_count" json:"ht_uid_count"`
	CopdUidCount  float64 `bson:"copd_uid_count" json:"copd_uid_count"`
	CaUidCount    float64 `bson:"ca_uid_count" json:"ca_uid_count"`
	PsyUidCount   float64 `bson:"psy_uid_count" json:"psy_uid_count"`
	HdCvdUidCount float64 `bson:"hd_cvd_uid_count" json:"hd_cvd_uid_count"`
}
