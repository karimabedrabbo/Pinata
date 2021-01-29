package handlers

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/apputils"
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"github.com/karimabedrabbo/eyo/api/models/response"
	"net/http"
)

func (h *RhEnv) RHGetUserProfile() (int, *dbmodel.User, error) {
	var err error

	userProfileGet := &requests.UserProfileGet{}
	if userProfileGet.UserId = h.Auth.GetUserId(h.C); userProfileGet.UserId != 0 {
		return http.StatusBadRequest, nil, apperror.RequestTokenIdentityMissing
	}

	var user *dbmodel.User
	if user, err = h.E.DBGetUserById(userProfileGet.UserId); err != nil {
		return http.StatusNotFound, nil, err
	}

	return http.StatusOK, user, nil
}

func (h *RhEnv) RHGetUserAvatar() (int, *response.SignedUrlResponse, error) {
	var err error
	userAvatarGet := &requests.UserAvatarGet{}

	if err = h.MapUri(userAvatarGet); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var attachments *[]dbmodel.Attachment
	if attachments, err = h.E.DBListAttachmentsByReference(
		k.AttachmentCategoryTypeUser,
		k.AttachmentMediaTypeImage,
		k.AttachmentUsedAsTypeAvatar,
		userAvatarGet.UserId,
		0,
		1,
	); err != nil {
		return http.StatusNotFound, nil, err
	}

	attachment := (*attachments)[0]

	var url *response.SignedUrlResponse
	if url, err = h.Storage.GetSignedUrl(
		http.MethodGet,
		attachment.MediaType,
		attachment.EncodingType,
		attachment.Md5,
		attachment.Id,
		attachment.TokenUuid); err != nil {
		return http.StatusInternalServerError, nil, err
	}


	return http.StatusOK, url, nil
}

func (h *RhEnv) RHPostUserProfile() (int, error) {
	var err error

	userPost := &requests.UserProfilePost{}
	var userId int64
	if userId = h.Auth.GetUserId(h.C); userId != 0 {
		return http.StatusBadRequest, apperror.RequestTokenIdentityMissing
	}
	userPost.UserId = userId

	if err = h.MapJson(userPost); err != nil {
		return http.StatusBadRequest, err
	}

	user := &dbmodel.User{}
	if err = user.PrepareUserPost(userPost); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBPostUserById(user); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}


func (h *RhEnv) RHPostUserAvatar() (int, *response.SignedUrlResponse, error) {
	var err error
	userAvatarPost := &requests.UserAvatarPost{}
	if err = h.MapUriJson(userAvatarPost); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if userAvatarPost.TokenUuid, err = apputils.GetUUID(); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	attachment := &dbmodel.Attachment{}
	if err = attachment.PrepareUserAvatarPost(userAvatarPost); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if err = h.E.DBPutAttachment(attachment); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var url *response.SignedUrlResponse
	if url, err = h.Storage.GetSignedUrl(
		http.MethodPut,
		attachment.MediaType,
		attachment.EncodingType,
		attachment.Md5,
		attachment.Id,
		attachment.TokenUuid); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, url, nil
}


