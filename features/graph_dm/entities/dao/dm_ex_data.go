package dao

type DmExpenseData struct {
	Year         int     `bson:"year" json:"year"`
	Month        int     `bson:"month" json:"month"`
	DmExpense    float64 `bson:"dm_expense" json:"dm_expense"`
	HgExpense    float64 `bson:"hg_expense" json:"hg_expense"`
	DmCkdExpense float64 `bson:"dm_ckd_expense" json:"dm_ckd_expense"`
	DmAcsExpense float64 `bson:"dm_acs_expense" json:"dm_acs_expense"`
	DmCvaExpense float64 `bson:"dm_cva_expense" json:"dm_cva_expense"`
}
