package handlers

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"github.com/karimabedrabbo/eyo/api/models/response"
	"net/http"
)

func (h *RhEnv) RHConfirmAttachment() (int, error) {
	var err error
	attachmentConfirmPost := &requests.AttachmentConfirmPost{}

	if err = h.MapJson(attachmentConfirmPost); err != nil {
		return http.StatusBadRequest, err
	}

	var attachment *dbmodel.Attachment
	if attachment, err = h.E.DBGetAttachmentById(attachmentConfirmPost.AttachmentId); err != nil {
		return http.StatusNotFound, err
	}

	if attachment.ConfirmedAt > 0 {
		return http.StatusUnauthorized, apperror.RequestAlreadyUsed
	}

	if err = attachment.PrepareAttachmentConfirmPost(attachmentConfirmPost); err != nil {
		return http.StatusInternalServerError, err
	}
	if err = h.E.DBPostAttachmentById(attachment); err != nil {
		return http.StatusInternalServerError, err
	}


	return http.StatusOK, nil
}

func (h *RhEnv) RHGetAttachment() (int, *response.SignedUrlResponse, error) {
	var err error
	attachmentGet := &requests.AttachmentGet{}

	if err = h.MapUri(attachmentGet); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var attachment *dbmodel.Attachment
	if attachment, err = h.E.DBGetAttachmentById(attachmentGet.AttachmentId); err != nil {
		return http.StatusNotFound, nil, err
	}

	var url *response.SignedUrlResponse
	if url, err = h.Storage.GetSignedUrl(http.MethodGet,
		attachment.MediaType,
		attachment.EncodingType,
		attachment.Md5,
		attachment.Id,
		attachment.TokenUuid); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, url, nil
}

func (h *RhEnv) RHGetAttachmentList() (int, *[]response.SignedUrlResponse, error) {
	var err error
	attachmentList := &requests.AttachmentList{}

	if err = h.MapQuery(attachmentList); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var attachments *[]dbmodel.Attachment
	if attachments, err = h.E.DBListAttachmentsByReference(
		attachmentList.CategoryType,
		attachmentList.MediaType,
		attachmentList.UsedAsType,
		attachmentList.ReferenceId,
		attachmentList.AfterId,
		attachmentList.Limit,
		); err != nil {
		return http.StatusNotFound, nil, err
	}

	urls := make([]response.SignedUrlResponse, len(*attachments))

	for i, attachment := range *attachments {
		var url *response.SignedUrlResponse
		if url, err = h.Storage.GetSignedUrl(http.MethodGet,
			attachment.MediaType,
			attachment.EncodingType,
			attachment.Md5,
			attachment.Id,
			attachment.TokenUuid); err != nil {
			return http.StatusInternalServerError, nil, err
		}

		urls[i] = *url
	}

	return http.StatusOK, &urls, nil

}

func (h *RhEnv) RHDeleteAttachment() (int, *response.SignedUrlResponse, error) {
	var err error
	attachmentDelete := &requests.AttachmentDelete{}

	if err = h.MapUri(attachmentDelete); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var attachment *dbmodel.Attachment
	if attachment, err = h.E.DBGetAttachmentById(attachmentDelete.AttachmentId); err != nil {
		return http.StatusNotFound, nil, err
	}

	if err = h.E.DBDeleteAttachmentById(attachmentDelete.AttachmentId); err != nil {
		return http.StatusNotFound, nil, err
	}

	var url *response.SignedUrlResponse
	if url, err = h.Storage.GetSignedUrl(http.MethodDelete,
		attachment.MediaType,
		attachment.EncodingType,
		attachment.Md5,
		attachment.Id,
		attachment.TokenUuid); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, url, nil
}


