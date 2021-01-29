package dbmodel

type Feedback struct {
	BaseModel
	AccountId int64 `gorm:"index;not null" json:"account_id"`
	Title        string `gorm:"size:255" json:"reason"`
	Comment       string `gorm:"text" json:"comment"`
}

