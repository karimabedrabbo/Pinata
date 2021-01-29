package requests


type ParticipantPut struct {
	ConversationId int64 `json:"conversation_id" binding:"required,numeric"`
}

type ParticipantPost struct {
	ConversationId int64 `json:"conversation_id" binding:"required,numeric"`
	Nickname string `json:"nickname" binding:"omitempty'"`
	Mute *bool `json:"mute" binding:"omitempty"`
}

type ParticipantGet struct {
}

type ParticipantList struct {
	ListRequest
	ConversationId int64 `json:"conversation_id" binding:"required,numeric"`
}

//leave or kick (if owner)
type ParticipantDelete struct {

}
