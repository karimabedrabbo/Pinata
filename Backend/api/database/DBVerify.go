package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)


func (e *DbEnv) DBGetVerifyById(verifyId int64) (*dbmodel.Verify, error) {
	var err error

	if verifyId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	verify := &dbmodel.Verify{}
	err = e.GetTx().Model(&dbmodel.Verify{}).First(verify, verifyId).Error
	if err != nil {
		return nil, err
	}

	return verify, nil
}

func (e *DbEnv) DBGetVerifyByAccountId(accountId int64) (*dbmodel.Verify, error) {
	var err error

	if accountId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	verify := &dbmodel.Verify{}
	err = e.GetTx().Model(&dbmodel.Verify{}).Where("account_id = ?", accountId).First(verify).Error
	if err != nil {
		return nil, err
	}

	return verify, nil
}

func (e *DbEnv) DBPutVerify(verify *dbmodel.Verify) error {
	var err error

	if verify == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(&dbmodel.Verify{}).Create(verify).Error
	if err != nil {
		return  err
	}

	return nil
}

func (e *DbEnv) DBPostVerifyById(verify *dbmodel.Verify) error {
	var err error

	if verify.Id == 0 {
		return apperror.DatabaseIdentityUninitialized
	}
	if verify == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(&dbmodel.Verify{}).Update(verify).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBCheckVerifiedByAccountId(accountId int64) (bool, error) {
	var err error

	if accountId == 0 {
		return false, apperror.DatabaseIdentityUninitialized
	}

	var verify *dbmodel.Verify
	if verify, err = e.DBGetVerifyByAccountId(accountId); err != nil {
		return false, err
	}

	return verify.UsedAt > 0, nil
}

func (e *DbEnv) DBDeleteVerificationsByAccountId(accountId int64) error {
	var err error

	if accountId == 0 {
		return apperror.DatabaseIdentityUninitialized
	}

	emptyVerify := &dbmodel.Verify{}
	err = e.GetTx().Model(emptyVerify).Where("account_id = ?", accountId).Delete(emptyVerify).Error
	if err != nil {
		return err
	}

	return nil
}