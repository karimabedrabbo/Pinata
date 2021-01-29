package database

import dbmodel "github.com/karimabedrabbo/eyo/api/models/database"

func (e *DbEnv) DBGetBanByAccountEmail(email string) (*dbmodel.Ban, error) {
	var err error

	var account *dbmodel.Account
	if account, err = e.DBGetAccountByEmail(email); err != nil {
		return nil, err
	}

	return e.DBGetBanByAccountId(account.Id)

}

func (e *DbEnv) DBCheckBannedByAccountEmail(email string) (bool, error) {
	var err error

	var account *dbmodel.Account
	if account, err = e.DBGetAccountByEmail(email); err != nil {
		return true, err
	}

	return e.DBCheckBannedByAccountId(account.Id)
}