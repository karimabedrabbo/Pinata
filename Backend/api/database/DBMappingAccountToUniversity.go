package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"strings"
)

func (e *DbEnv) DBGetUniversityByAccountEmail(email string) (*dbmodel.University, error) {
	if email == "" {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	components := strings.Split(email, "@")
	if len(components) != 2 {
		return nil, apperror.DatabaseParsingError
	}
	domain := components[1]

	return e.DBGetUniversityByDomain(domain)
}


func (e *DbEnv) DBGetUniversityByAccountId(accountId int64) (*dbmodel.University, error) {
	var err error

	if accountId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	var account *dbmodel.Account
	if account, err = e.DBGetAccountById(accountId); err != nil {
		return nil, err
	}

	return e.DBGetUniversityByAccountEmail(account.Email)
}

ck4r3j6zq0055a4p8mlmwt6k8
