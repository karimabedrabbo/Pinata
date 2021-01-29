package handlers

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"net/http"
)

func (h *RhEnv) RHCloseAccount() (int, error) {
	var err error


	//this request model is for consistency (trivially you could just pass the fn id parameter into the database)
	closeAccountPost := &requests.CloseAccountPost{}
	var accountId int64
	if accountId = h.Auth.GetAccountId(h.C); accountId == 0 {
		return http.StatusBadRequest, apperror.RequestTokenIdentityMissing
	}
	closeAccountPost.AccountId = accountId

	if err = h.E.DBDeleteAccountById(closeAccountPost.AccountId); err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}

