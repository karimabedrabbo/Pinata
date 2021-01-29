package requests

import "github.com/google/uuid"

type PasswordResetPut struct {
	AccountId int64 `json:"-" binding:"-"`
	HashedTokenUuid string `json:"-" binding:"-"`
	Email string `json:"email" binding:"required,email,endswith=.edu"`
}

type PasswordResetPost struct {
	AccountId int64 `json:"-" binding:"-"`
	PasswordResetId int64 `form:"password_reset_id" json:"password_reset_id" binding:"required,numeric"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
	TokenUuid uuid.UUID `form:"token_uuid" json:"token_uuid" binding:"required,uuid"`
}
