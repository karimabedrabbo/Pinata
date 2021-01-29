package handlers

import (
	"github.com/karimabedrabbo/eyo/api/apputils"
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"github.com/karimabedrabbo/eyo/api/models/response"
	"net/http"
)

func (h *RhEnv) RHGetUniversity() (int, *dbmodel.University, error) {
	var err error
	univGet := &requests.UniversityGet{}

	if err = h.MapUri(univGet); err != nil {
		return http.StatusBadRequest, nil, err
	}


	var university *dbmodel.University
	if university, err = h.E.DBGetUniversityById(univGet.UniversityId); err != nil {
		return http.StatusNotFound, nil, err
	}


	return http.StatusOK, university, nil

}

func (h *RhEnv) RHGetUniversityAvatar() (int, *response.SignedUrlResponse, error) {
	var err error
	univAvatarGet := &requests.UniversityAvatarGet{}

	if err = h.MapUri(univAvatarGet); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var attachments *[]dbmodel.Attachment
	if attachments, err = h.E.DBListAttachmentsByReference(
		k.AttachmentCategoryTypeUniversity,
		k.AttachmentMediaTypeImage,
		k.AttachmentUsedAsTypeAvatar,
		univAvatarGet.UniversityId,
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


func (h *RhEnv) RHGetUniversityList() (int, *[]dbmodel.University, error) {
	var err error

	univList := &requests.UniversityList{}
	if err = h.MapQuery(univList); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var univs *[]dbmodel.University
	if univs, err = h.E.DBListUniversities(univList.AfterId, univList.Limit); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, univs, nil
}


func (h *RhEnv) RHPutUniversity() (int, error) {
	var err error

	univPut := &requests.UniversityPut{}
	if err = h.MapJson(univPut); err != nil {
		return http.StatusBadRequest, err
	}

	univ := &dbmodel.University{}
	if err = univ.PrepareUniversityPut(univPut); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBPutUniversity(univ); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}


func (h *RhEnv) RHPostUniversityAvatar() (int, *response.SignedUrlResponse, error) {
	var err error
	univAvatarPost := &requests.UniversityAvatarPost{}
	if err = h.MapUriJson(univAvatarPost); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if univAvatarPost.TokenUuid, err = apputils.GetUUID(); err != nil {
		return http.StatusInternalServerError, nil, err
	}


	attachment := &dbmodel.Attachment{}
	if err = attachment.PrepareUniversityAvatarPost(univAvatarPost); err != nil {
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


func (h *RhEnv) RHPostUniversity() (int, error) {
	var err error

	univPost := &requests.UniversityPost{}
	if err = h.MapJson(univPost); err != nil {
		return http.StatusBadRequest, err
	}

	univ := &dbmodel.University{}
	if err = univ.PrepareUniversityPost(univPost); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBPostUniversityById(univ); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (h *RhEnv) RHDeleteUniversity() (int, error) {
	var err error

	univDelete := &requests.UniversityDelete{}
	if err = h.MapUri(univDelete); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBDeleteUniversityById(univDelete.UniversityId); err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
