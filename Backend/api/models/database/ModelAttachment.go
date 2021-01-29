package dbmodel

import (
	"github.com/google/uuid"
	"github.com/karimabedrabbo/eyo/api/apperror"
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"time"
)

type Attachment struct {
	BaseModel
	Md5      string `gorm:"unique;size:100;not null" json:"md5"`
	CategoryType string `gorm:"size:100;not null;index" json:"category_type"`
	MediaType string `gorm:"size:100;not null" json:"media_type"`
	EncodingType string `gorm:"size:100;not null" json:"encoding_type"`
	UsedAsType string `gorm:"size:100;not null" json:"used_as_type"`
	UniversityId   int64  `gorm:"index" json:"university_id"`
	UserId         int64  `gorm:"index" json:"user_id"`
	MessageId      int64  `gorm:"index" json:"message_id"`
	ConversationId  int64  `gorm:"index" json:"conversation_id"`
	TokenUuid uuid.UUID `gorm:"type:uuid;not null" json:"token_uuid"`
	ConfirmedAt    int64 `json:"confirmed_at"`
}

func (a *Attachment) helperPrepareAttachmentPost(tokenUuid uuid.UUID, md5 string, encodingType string) error {
	if tokenUuid == uuid.Nil {
		return apperror.PrepareIdentityUninitialized
	}
	if md5 == "" {
		return apperror.PrepareMissingAttribute
	}
	if encodingType == "" {
		return apperror.PrepareMissingAttribute
	}
	a.TokenUuid = tokenUuid
	a.Md5 = md5
	a.EncodingType = encodingType
	return nil
}

func (a *Attachment) PrepareImageAttachmentPost(r *requests.ImageAttachmentCreatable) error {
	return a.helperPrepareAttachmentPost(r.TokenUuid, r.Md5, r.EncodingType)
}

func (a *Attachment) PrepareVideoAttachmentPost(r *requests.VideoAttachmentCreatable) error {
	return a.helperPrepareAttachmentPost(r.TokenUuid, r.Md5, r.EncodingType)
}


func (a *Attachment) PrepareAttachmentConfirmPost(r *requests.AttachmentConfirmPost) error {
	if r.AttachmentId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	if r.TokenUuid == uuid.Nil {
		return apperror.PrepareMissingAttribute
	}
	a.Id = r.AttachmentId
	a.TokenUuid = r.TokenUuid
	a.ConfirmedAt = time.Now().Unix()
	a.Prepare()
	return nil
}

func (a *Attachment) PrepareUniversityAvatarPost(r *requests.UniversityAvatarPost) error {
	var err error
	if r.UniversityId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	a.UniversityId = r.UniversityId
	a.MediaType = k.AttachmentMediaTypeImage
	a.CategoryType = k.AttachmentCategoryTypeUniversity
	a.UsedAsType = k.AttachmentUsedAsTypeAvatar
	if err = a.PrepareImageAttachmentPost(r.ImageAttachmentCreatable); err != nil {
		return err
	}
	a.Prepare()
	return nil
}

func (a *Attachment) PrepareUserAvatarPost(r *requests.UserAvatarPost) error {
	var err error
	if r.UserId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	a.UserId = r.UserId
	a.MediaType = k.AttachmentMediaTypeImage
	a.CategoryType = k.AttachmentCategoryTypeUser
	a.UsedAsType = k.AttachmentUsedAsTypeAvatar
	if err = a.PrepareImageAttachmentPost(r.ImageAttachmentCreatable); err != nil {
		return err
	}
	a.Prepare()
	return nil
}