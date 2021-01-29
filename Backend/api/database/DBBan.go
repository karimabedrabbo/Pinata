package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetBanByAccountId(accountId int64) (*dbmodel.Ban, error) {
	var err error

	if accountId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	account := &dbmodel.Account{}
	if account, err = e.DBGetAccountById(accountId); err != nil {
		return nil, err
	}

	ban := &dbmodel.Ban{}
	err = e.GetTx().Model(&dbmodel.Ban{}).Where("to_account_id = ?", account.Id).First(ban).Error
	if err != nil {
		return nil, err
	}

	return ban, nil
}


func (e *DbEnv) DBPutBan(ban *dbmodel.Ban) error {
	var err error

	if ban == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(&dbmodel.Ban{}).Create(ban).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBCheckBannedByAccountId(accountId int64) (bool, error) {
	var err error

	var ban *dbmodel.Ban
	if ban, err = e.DBGetBanByAccountId(accountId); err != nil {
		if e.IsRecordNotFound(err) {
			return false, nil
		}
		return true, err
	}

	return ban != nil, nil
}


func (e *DbEnv) DBDeleteBanByAccountId(accountId int64) error {
	var err error

	if accountId == 0 {
		return apperror.DatabaseIdentityUninitialized
	}

	err = e.GetTx().Model(&dbmodel.Ban{}).Where("to_account_id = ?", accountId).Delete(&dbmodel.Ban{}).Error
	if err != nil {
		return err
	}

	return nil
}
