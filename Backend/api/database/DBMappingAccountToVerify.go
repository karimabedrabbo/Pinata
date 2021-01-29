package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetVerifyByAccountEmail(email string) (*dbmodel.Verify, error) {
	var err error

	if email == "" {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	var account *dbmodel.Account
	if account, err = e.DBGetAccountByEmail(email); err != nil {
		return nil, err
	}

	return e.DBGetVerifyByAccountId(account.Id)
}

func (e *DbEnv) DBCheckVerifiedByAccountEmail(email string) (bool, error) {
	var err error

	if email == "" {
		return false, apperror.DatabaseIdentityUninitialized
	}

	var account *dbmodel.Account
	if account, err = e.DBGetAccountByEmail(email); err != nil {
		return false, err
	}

	return e.DBCheckVerifiedByAccountId(account.Id)
}

