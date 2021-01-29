package handlers

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/apputils"
	k "github.com/karimabedrabbo/eyo/api/constants"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"github.com/karimabedrabbo/eyo/api/models/response"
)

func (h *RhEnv) RHGinJwtIdentityAuthenticator() (*response.PayloadClaims, error) {
	var err error

	loginPost := &requests.LoginPost{}
	if err = h.MapJson(loginPost); err != nil {
		return nil, err
	}

	var account *dbmodel.Account
	if account, err = h.E.DBGetAccountByEmail(loginPost.Email); err != nil {
		return nil, err
	}

	if err = apputils.VerifyPassword(account.PasswordHash, loginPost.Password); err != nil {
		return nil, err
	}

	//I don't think this is possible (since signed up users get "user" role) but just in case
	if account.Role == k.AccountRoleAnonymous {
		return nil, apperror.AuthenticationUserUnverified
	}


	var isBanned bool
	if isBanned, err = h.E.DBCheckBannedByAccountId(account.Id); err != nil {
		return nil, err
	}
	if isBanned {
		return nil, apperror.AuthenticationUserBanned
	}

	var isVerified bool
	if isVerified, err = h.E.DBCheckVerifiedByAccountId(account.Id); err != nil {
		return nil, err
	}
	if !isVerified {
		return nil, apperror.AuthenticationUserUnverified
	}


	var university *dbmodel.University
	if university, err = h.E.DBGetUniversityByAccountEmail(loginPost.Email); err != nil {
		return nil, err
	}


	return &response.PayloadClaims{
		UserId: account.UserId,
		AccountId: account.Id,
		Role: account.Role,
		UniversityId: university.Id,
		UniversityVerified: true,
		EmailVerified: true,
	}, nil
}

func (h *RhEnv) RHGinJwtAnonymousAuthenticator() (*response.PayloadClaims, error) {
	var err error

	loginAnonPost := &requests.LoginAnonymousPost{}
	if err = h.MapJson(loginAnonPost); err != nil {
		return nil, err
	}

	//todo
	//check university exists

	user := &dbmodel.User{}
	if err = user.PrepareLoginAnonymousPost(loginAnonPost); err != nil {
		return nil, err
	}
	if err = h.E.DBPutUser(user); err != nil {
		return nil, err
	}

	return &response.PayloadClaims{
		UserId: user.Id,
		AccountId: 0,
		Role: k.AccountRoleAnonymous,
		UniversityId: loginAnonPost.UniversityId,
		UniversityVerified: false,
		EmailVerified: false,
	}, nil
}