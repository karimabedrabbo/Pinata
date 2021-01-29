package handlers

import (
	"github.com/google/uuid"
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/apputils"
	k "github.com/karimabedrabbo/eyo/api/constants"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"net/http"
	"time"
)

func (h *RhEnv) RHRequestVerifyAccount() (int, error) {
	var err error

	verifyPut := &requests.VerifyPut{}
	if err = h.MapJson(verifyPut); err != nil {
		return http.StatusBadRequest, err
	}

	var account *dbmodel.Account
	if account, err = h.E.DBGetAccountByEmail(verifyPut.Email); err != nil {
		return http.StatusNotFound, err
	}

	if err = h.E.DBDeleteVerificationsByAccountId(account.Id); err != nil {
		return http.StatusInternalServerError, err
	}

	var tokenUuid uuid.UUID
	var tokenHashedUuid string
	if tokenUuid, err = apputils.GetUUID(); err != nil {
		return http.StatusInternalServerError, err
	}
	if tokenHashedUuid, err = apputils.DerivePasswordHash(tokenUuid.String()); err != nil {
		return http.StatusInternalServerError, err
	}

	verifyPut.AccountId = account.Id
	verifyPut.HashedTokenUuid = tokenHashedUuid

	verify := &dbmodel.Verify{}
	if err = verify.PrepareVerifyPut(verifyPut); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBPutVerify(verify); err != nil {
		return http.StatusInternalServerError, err
	}

	var user *dbmodel.User
	if user, err = h.E.DBGetUserById(account.UserId); err != nil {
		return http.StatusNotFound, err
	}

	//if the name for the user is blank, use their email instead
	if user.Name == "" {
		user.Name = account.Email
	}


	//use the frontend url here (not an api link)
	verifyUrl := apputils.GetAppUrl() + k.RouteVerifyAccount + "?verify_id=" + apputils.IdToString(verify.Id) + "&token_uuid=" + tokenUuid.String()


	if err = h.Mail.SendVerifyEmail(user.Name, account.Email, verifyUrl); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (h *RhEnv) RHConfirmVerifyAccount() (int, error) {
	var err error

	verifyPost := &requests.VerifyPost{}
	if err = h.MapJson(verifyPost); err != nil {
		return http.StatusBadRequest, err
	}

	var verify *dbmodel.Verify
	if verify, err = h.E.DBGetVerifyById(verifyPost.VerifyId); err != nil {
		return http.StatusNotFound, err
	}

	if err = apputils.VerifyPassword(verify.HashedTokenUuid, verifyPost.TokenUuid.String()); err != nil {
		return http.StatusUnauthorized, apperror.RequestInvalidCredentials
	}

	if verify.ExpiresAt > time.Now().Unix() {
		return http.StatusUnauthorized, apperror.RequestWindowExpired
	}

	if verify.UsedAt > 0 {
		return http.StatusUnauthorized, apperror.RequestAlreadyUsed
	}

	if err = verify.PrepareVerifyPost(verifyPost); err != nil {
		return http.StatusInternalServerError, err
	}
	if err = h.E.DBPostVerifyById(verify); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

