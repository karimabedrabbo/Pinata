package requests

import "github.com/google/uuid"

type VerifyPut struct {
	AccountId int64 `json:"-" binding:"-"`
	HashedTokenUuid string `json:"-" binding:"-"`
	Email string `json:"email" binding:"required,email,endswith=.edu"`
}

type VerifyPost struct {
	VerifyId int64 `form:"verify_id" json:"verify_id" binding:"required,numeric"`
	TokenUuid uuid.UUID `form:"token_uuid" json:"token_uuid" binding:"required,uuid"`
}
