package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	k "github.com/karimabedrabbo/eyo/api/constants"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetAttachmentById(attachmentId int64) (*dbmodel.Attachment, error) {
	var err error

	if attachmentId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	a := &dbmodel.Attachment{}
	if err = e.GetTx().Model(&dbmodel.Attachment{}).First(a, attachmentId).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (e *DbEnv) DBListAttachmentsByReference(categoryType string, mediaType string, usedAsType string, referenceId int64, afterId int64, limit int64) (*[]dbmodel.Attachment, error) {
	var err error

	temp := make([]dbmodel.Attachment, 0)
	attachments := &temp


	if categoryType != "" {
		e.SetTx(e.GetTx().Where("category_type = ?", categoryType))
	}

	if referenceId != 0 {
		switch categoryType {
		case k.AttachmentCategoryTypeUniversity:
			e.SetTx(e.GetTx().Where("university_id = ?", referenceId))
		case k.AttachmentCategoryTypeUser:
			e.SetTx(e.GetTx().Where("user_id = ?", referenceId))
		case k.AttachmentCategoryTypeConversation:
			e.SetTx(e.GetTx().Where("conversation_id = ?", categoryType))
		}
	}

	if mediaType != "" {
		e.SetTx(e.GetTx().Where("media_type = ?", mediaType))
	}

	if usedAsType != "" {
		e.SetTx(e.GetTx().Where("used_as_type = ?", usedAsType))
	}

	e.DBListRequest("created_at desc", afterId, limit)

	err = e.GetTx().Model(&dbmodel.Attachment{}).Find(attachments).Error
	if err != nil {
		return nil, err
	}

	return attachments, nil
}

func (e *DbEnv) DBPostAttachmentById(attachment *dbmodel.Attachment) error {
	var err error

	if attachment.Id == 0 {
		return apperror.DatabaseModelUninitialized
	}
	if attachment == nil {
		return apperror.DatabaseModelUninitialized
	}

	if err = e.GetTx().Model(&dbmodel.Attachment{}).Update(attachment).Error; err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBPutAttachment(attachment *dbmodel.Attachment) error {
	var err error

	if attachment == nil {
		return apperror.DatabaseModelUninitialized
	}
	if err = e.GetTx().Model(&dbmodel.Attachment{}).Create(attachment).Error; err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBDeleteAttachmentById(attachmentId int64) error {
	var err error

	if attachmentId == 0 {
		return apperror.DatabaseIdentityUninitialized
	}

	if err = e.GetTx().Model(&dbmodel.Attachment{}).Delete(attachmentId).Error; err != nil {
		return err
	}
	return nil
}
