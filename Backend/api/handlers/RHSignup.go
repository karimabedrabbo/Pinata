package handlers

import (
	"github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"net/http"
)

func (h *RhEnv) RHSignupUser() (int, error) {
	var err error

	signupPut := &requests.SignupPut{}
	if err =  h.MapJson(signupPut); err != nil {
		return http.StatusBadRequest, err
	}

	if _, err = h.E.DBGetUniversityByAccountEmail(signupPut.Email); err != nil {
		//no matching universities
		return http.StatusBadRequest, err
	}

	account := &dbmodel.Account{}
	if err = account.PrepareSignupPut(signupPut); err != nil {
		return http.StatusBadRequest, err
	}
	if err = h.E.DBPutAccount(account); err != nil {
		return http.StatusBadRequest, err
	}

	user := &dbmodel.User{}
	if err = user.PrepareSignupPut(signupPut); err != nil {
		return http.StatusBadRequest, err
	}
	if err = h.E.DBPutUser(user); err != nil {
		return http.StatusInternalServerError, err
	}


	return http.StatusOK, nil
}
