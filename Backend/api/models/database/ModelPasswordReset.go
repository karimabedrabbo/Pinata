package dbmodel

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"time"
)

type PasswordReset struct {
	BaseModel
	AccountId int64 `gorm:"not null" json:"account_id"`
	HashedTokenUuid string `gorm:"size:100;not null" json:"hashed_token_uuid"`
	UsedAt int64 `json:"used_at"`
	ExpiresAt int64 `gorm:"not null" json:"expires_at"`
}


func (reset *PasswordReset) PreparePasswordResetPut(r *requests.PasswordResetPut) error {
	if r.AccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	if r.HashedTokenUuid == "" {
		return apperror.PrepareMissingAttribute
	}
	reset.AccountId = r.AccountId
	reset.HashedTokenUuid = r.HashedTokenUuid
	reset.ExpiresAt = time.Now().Add(time.Hour).Unix()
	reset.Prepare()
	return nil
}


func (reset *PasswordReset) PreparePasswordResetPost(r *requests.PasswordResetPost) error {
	if r.PasswordResetId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	reset.Id = r.PasswordResetId
	reset.UsedAt = time.Now().Unix()
	reset.Prepare()
	return nil
}