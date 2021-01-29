package dbmodel

type Block struct {
	BaseModel
	FromAccountId int64 `gorm:"index" json:"from_account_id"`
	ToAccountId   int64 `gorm:"index" json:"to_account_id"`
}
