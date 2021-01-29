package dbmodel

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"time"
)

type Verify struct {
	BaseModel
	AccountId int64 `gorm:"not null;index" json:"account_id"` //references account
	HashedTokenUuid string `gorm:"size:100;not null" json:"hashed_token_uuid"`
	UsedAt int64 `json:"used_at"`
	ExpiresAt int64 `gorm:"not null" json:"expires_at"`
}

func (verify *Verify) PrepareVerifyPut(r *requests.VerifyPut) error {
	if r.AccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	if r.HashedTokenUuid == "" {
		return apperror.PrepareMissingAttribute
	}
	verify.AccountId = r.AccountId
	verify.HashedTokenUuid = r.HashedTokenUuid
	verify.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	verify.Prepare()
	return nil
}


func (verify *Verify) PrepareVerifyPost(r *requests.VerifyPost) error {
	if r.VerifyId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	verify.Id = r.VerifyId
	verify.UsedAt = time.Now().Unix()
	verify.Prepare()
	return nil
}
