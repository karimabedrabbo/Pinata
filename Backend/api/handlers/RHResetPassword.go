package handlers

import (
	"github.com/google/uuid"
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/apputils"
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"net/http"
	"time"
)

func (h *RhEnv) RHRequestPasswordReset() (int, error) {
	var err error

	resetPut := &requests.PasswordResetPut{}
	if err = h.MapJson(resetPut); err != nil {
		return http.StatusBadRequest, err
	}

	var account *dbmodel.Account
	if account, err = h.E.DBGetAccountByEmail(resetPut.Email); err != nil {
		return http.StatusNotFound, err
	}

	if err = h.E.DBDeletePasswordResetsByAccountId(account.Id); err != nil {
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

	resetPut.AccountId = account.Id
	resetPut.HashedTokenUuid = tokenHashedUuid

	reset := &dbmodel.PasswordReset{}
	if err = reset.PreparePasswordResetPut(resetPut); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBPutPasswordReset(reset); err != nil {
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
	resetUrl := apputils.GetAppUrl() + k.RoutePasswordReset + "?password_reset_id=" + apputils.IdToString(reset.Id) + "&token_uuid=" + tokenUuid.String()

	if err = h.Mail.SendPasswordReset(user.Name, account.Email, resetUrl); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (h *RhEnv) RHConfirmPasswordReset() (int, error) {
	var err error

	resetPost := &requests.PasswordResetPost{}
	if err = h.MapJson(resetPost); err != nil {
		return http.StatusBadRequest, err
	}

	var reset *dbmodel.PasswordReset
	if reset, err = h.E.DBGetPasswordResetById(resetPost.PasswordResetId); err != nil {
		return http.StatusNotFound, err
	}

	if err = apputils.VerifyPassword(reset.HashedTokenUuid, resetPost.TokenUuid.String()); err != nil {
		return http.StatusUnauthorized, apperror.RequestInvalidCredentials
	}

	if reset.ExpiresAt > time.Now().Unix() {
		return http.StatusUnauthorized, apperror.RequestWindowExpired
	}

	if reset.UsedAt > 0 {
		return http.StatusUnauthorized, apperror.RequestAlreadyUsed
	}

	if err = reset.PreparePasswordResetPost(resetPost); err != nil {
		return http.StatusInternalServerError, err
	}
	if err = h.E.DBPostPasswordResetById(reset); err != nil {
		return http.StatusInternalServerError, err
	}


	resetPost.AccountId = reset.AccountId
	account := &dbmodel.Account{}
	if err = account.PreparePasswordResetPost(resetPost); err != nil {
		return http.StatusInternalServerError, err
	}
	if err = h.E.DBPostAccountById(account); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

