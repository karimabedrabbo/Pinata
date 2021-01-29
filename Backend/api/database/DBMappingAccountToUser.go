package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetUserByAccountId(accountId int64) (*dbmodel.User, error){
	var err error

	if accountId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	var account *dbmodel.Account
	if account, err = e.DBGetAccountById(accountId); err != nil {
		return nil, err
	}

	return e.DBGetUserById(account.UserId)
}
