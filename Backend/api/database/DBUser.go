package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetUserById(userId int64) (*dbmodel.User, error) {
	var err error

	if userId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	u := &dbmodel.User{}
	if err = e.GetTx().Model(&dbmodel.User{}).First(u, userId).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (e *DbEnv) DBPutUser(user *dbmodel.User) error {
	var err error

	if user == nil {
		return apperror.DatabaseIdentityUninitialized
	}

	err = e.GetTx().Model(&dbmodel.User{}).Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBPostUserById(user *dbmodel.User) error {
	var err error

	if user.Id == 0 {
		return apperror.DatabaseIdentityUninitialized
	}
	if user == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(user).Update(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBDeleteUserById(userId int64) error {
	var err error

	if userId == 0 {
		return apperror.DatabaseIdentityUninitialized
	}

	if err = e.GetTx().Model(&dbmodel.User{}).Delete(userId).Error; err != nil {
		return err
	}
	return nil
}
