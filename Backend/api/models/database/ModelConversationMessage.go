package dbmodel

type ConversationMessage struct {
	BaseModel
	TransactionSeqNum         int64       `gorm:"not null" json:"transaction_seq_num"`
	ConversationId           int64        `gorm:"index;not null" json:"conversation_id"`
	ParticipantId            int64        `gorm:"index;not null" json:"participant_id"`
	Participant              *ConversationParticipant `gorm:"foreignkey:ParticipantId"`
	MessageType              string       `gorm:"size:100" json:"message_type"`
	Status                   string       `gorm:"size:100" json:"status"`
	Content                  string       `gorm:"text;not null" json:"content"`
	ResponseId               int64        `gorm:"index" json:"response_id"`
	Response                 *ConversationMessage     `gorm:"foreignkey:ResponseId"`
	IsMutable                bool        `gorm:"default:false" json:"is_mutable"`
	IsRespondable            bool        `gorm:"default:true" json:"is_respondable"`
}

