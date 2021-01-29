package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetAccountById(accountId int64) (*dbmodel.Account, error) {
	var err error

	if accountId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	a := &dbmodel.Account{}
	if err = e.GetTx().Model(&dbmodel.Account{}).First(a, accountId).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (e *DbEnv) DBGetAccountByEmail(email string) (*dbmodel.Account, error) {
	var err error

	if email == "" {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	a := &dbmodel.Account{}
	if err = e.GetTx().Model(&dbmodel.Account{}).Where("email = ?", email).First(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

//user might not have account
func (e *DbEnv) DBGetAccountByUserId(userId int64) (*dbmodel.Account, error){
	var err error

	if userId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	a := &dbmodel.Account{}
	if err = e.GetTx().Model(&dbmodel.Account{}).Where("user_id = ?", userId).First(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (e *DbEnv) DBPutAccount(account *dbmodel.Account) error {
	var err error

	if account == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(&dbmodel.Account{}).Create(account).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBPostAccountById(account *dbmodel.Account) error {
	var err error

	if account.Id == 0 {
		return apperror.DatabaseIdentityUninitialized
	}
	if account == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(account).Update(account).Error
	if err != nil {
		return err
	}

	return nil
}


func (e *DbEnv) DBDeleteAccountById(accountId int64) error {
	var err error

	if accountId == 0 {
		return apperror.DatabaseIdentityUninitialized
	}

	if err = e.GetTx().Model(&dbmodel.Account{}).Delete(accountId).Error; err != nil {
		return err
	}
	return nil
}


