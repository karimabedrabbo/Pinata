package requests

import "github.com/google/uuid"

type AttachmentGet struct {
	AttachmentId int64 `uri:"attachment_id" json:"attachment_id" binding:"required,numeric"`
}

type AttachmentConfirmPost struct {
	//todo depends on what google returns
	AttachmentId int64 `json:"attachment_id" binding:"required,numeric"`
	TokenUuid uuid.UUID `json:"token_uuid" binding:"required,uuid"`
}

type AttachmentList struct {
	ListRequest
	ReferenceId int64 `form:"reference_id" json:"reference_id" binding:"omitempty,numeric"`
	MediaType string `form:"media_type" json:"media_type" binding:"omitempty,oneof=image video"`
	UsedAsType	string `form:"used_as_type" json:"used_as_type" binding:"omitempty,oneof=avatar message"`
	CategoryType string `form:"category_type" json:"category_type" binding:"omitempty,oneof=university conversation user"`
}

type AttachmentDelete struct {
	AttachmentId int64 `uri:"attachment_id" json:"attachment_id" binding:"required,numeric"`
}
