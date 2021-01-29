package requests

type ConversationPut struct {
	ChannelId  int64 `json:"channel_id" binding:"required_without=UserId,numeric"`
	UserId  int64 `json:"channel_id" binding:"required_without=ChannelId,numeric"`
}

type ConversationPost struct {
	Title string `json:"title" binding:"required"`
}

type ConversationGet struct {

}

type ConversationList struct {
	ListRequest
}

type ConversationDelete struct {

}


type conversationAttachmentBody struct {
	UserId         int64  `json:"user_id" binding:"required,numeric"`
	ConversationId int64  `json:"conversation_id" binding:"required,numeric"`
	MessageId      int64  `json:"message_id" binding:"required,numeric"`
}
