package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetPasswordResetById(passwordResetId int64) (*dbmodel.PasswordReset, error) {
	var err error

	if passwordResetId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	reset := &dbmodel.PasswordReset{}
	err = e.GetTx().Model(&dbmodel.PasswordReset{}).First(reset, passwordResetId).Error
	if err != nil {
		return nil, err
	}

	return reset, nil
}

func (e *DbEnv) DBPutPasswordReset(reset *dbmodel.PasswordReset) error {
	var err error

	if reset == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(&dbmodel.PasswordReset{}).Create(reset).Error
	if err != nil {
		return  err
	}

	return nil
}

func (e *DbEnv) DBPostPasswordResetById(reset *dbmodel.PasswordReset) error {
	var err error

	if reset.Id == 0 {
		return apperror.DatabaseIdentityUninitialized
	}
	if reset == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(&dbmodel.PasswordReset{}).Update(reset).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBDeletePasswordResetsByAccountId(accountId int64) error {
	var err error

	if accountId == 0 {
		return apperror.DatabaseIdentityUninitialized
	}

	emptyReset := &dbmodel.PasswordReset{}
	err = e.GetTx().Model(emptyReset).Where("account_id = ?", accountId).Delete(emptyReset).Error
	if err != nil {
		return err
	}

	return nil
}

