package handlers

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"net/http"
)

func (h *RhEnv) RHGetBan() (int, *dbmodel.Ban, error) {
	var err error

	var accountId int64
	if accountId = h.Auth.GetAccountId(h.C); accountId != 0 {
		return http.StatusBadRequest, nil, apperror.RequestTokenIdentityMissing
	}

	var ban *dbmodel.Ban
	if ban, err = h.E.DBGetBanByAccountId(accountId); err != nil {
		return http.StatusNotFound, nil, err
	}

	return http.StatusOK, ban, nil
}

func (h *RhEnv) RHPutBan() (int, error) {
	var err error

	var accountId int64
	if accountId = h.Auth.GetAccountId(h.C); accountId != 0 {
		return http.StatusBadRequest, apperror.RequestTokenIdentityMissing
	}

	banPut := &requests.BanPut{}
	if err = h.MapJson(banPut); err != nil {
		return http.StatusBadRequest, err
	}

	ban := &dbmodel.Ban{}
	if err = ban.PrepareBanPut(banPut); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBPutBan(ban); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (h *RhEnv) RHDeleteBan() (int, error) {
	var err error

	banDelete := &requests.BanDelete{}
	if err = h.MapUri(banDelete); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBDeleteUniversityById(banDelete.ToAccountId); err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}