package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetUniversityById(univId int64) (*dbmodel.University, error) {
	var err error
	if univId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	univ := &dbmodel.University{}

	err = e.GetTx().Model(&dbmodel.University{}).First(univ, univId).Error
	if err != nil {
		return nil, err
	}

	return univ, nil
}

func (e *DbEnv) DBGetUniversityByDomain(domain string) (*dbmodel.University, error) {
	var err error
	if domain == "" {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	univ := &dbmodel.University{}
	err = e.GetTx().Model(&dbmodel.University{}).Where("domain = ?", domain).First(univ).Error
	if err != nil {
		return nil, err
	}

	return univ, nil
}


func (e *DbEnv) DBListUniversities(afterId int64, limit int64) (*[]dbmodel.University, error) {
	var err error

	temp := make([]dbmodel.University, 0)
	univs := &temp


	e.DBListRequest("name asc", afterId, limit)

	err = e.GetTx().Model(&dbmodel.University{}).Find(univs).Error
	if err != nil {
		return nil, err
	}

	return univs, nil
}

func (e *DbEnv) DBPutUniversity(univ *dbmodel.University) error {
	var err error

	if univ == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(&dbmodel.University{}).Create(univ).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBPostUniversityById(univ *dbmodel.University) error {
	var err error

	if univ.Id == 0 {
		return apperror.DatabaseIdentityUninitialized
	}
	if univ == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(univ).Update(univ).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBDeleteUniversityById(univId int64) error {
	var err error

	if univId == 0 {
		return apperror.DatabaseIdentityUninitialized
	}

	if err = e.GetTx().Unscoped().Model(&dbmodel.University{}).Delete(univId).Error; err != nil {
		return err
	}
	return nil
}

