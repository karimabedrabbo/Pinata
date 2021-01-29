package dbmodel

type Conversation struct {
	BaseModel
	Title        string  `gorm:"size:255" json:"title"`
	ChannelId int64 `gorm:"index" json:"channel_id"`
	//Participants *[]ConversationParticipant `gorm:"foreignkey:ParticipantId"`
	//Messages     *[]ConversationMessage `gorm:"foreignkey:MessageId"`
	//Attachments  *[]ConversationAttachment `gorm:"foreignkey:AttachmentId"`
	AvatarId   	int64  `gorm:"index" json:"avatar_id"`
}
