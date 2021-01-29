package dbmodel

type ConversationParticipant struct {
	BaseModel
	ConversationId int64 `gorm:"index" json:"conversation_id"`
	UserId      int64 `gorm:"index" json:"user_id"`
	Nickname       string `gorm:"size:255" json:"nickname"`
	Role 		string `gorm:"size:100" json:"role"`
	Muted        bool `gorm:"default:false" json:"muted"`
	Active       bool `gorm:"default:true" json:"active"`
}

