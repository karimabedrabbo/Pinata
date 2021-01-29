package requests


type MessagePut struct {
	TransactionSeqNum int64 `json:"transaction_seq_num" binding:"required,numeric"`
	ConversationId  int64 `json:"conversation_id" binding:"required,numeric"`
	ParticipantId   int64 `json:"participant_id" binding:"required,numeric"`
	ResponseId    	int64 `json:"response_id" binding:"omitempty,numeric"`
	MessageType string `json:"message_type" binding:"required,oneof=content ack reaction multi_attachment attachment poll member_add member_leave member_invite member_nickname_change conversation_title_change conversation_avatar_change"`
	Content         string `json:"content" binding:"omitempty"`
}

type MessagePost struct {

	Status string `json:"status" binding:"omitempty,oneof=queued processing delivered faliure"`
	Content string `json:"content" binding:"omitempty"`
}

type MessageGet struct {

}

type MessageList struct {
	ListRequest
	ConversationId int64 `json:"conversation_id" binding:"omitempty,numeric"`
}


