package database

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func (e *DbEnv) DBGetReportById(reportId int64) (*dbmodel.Report, error) {
	var err error

	if reportId == 0 {
		return nil, apperror.DatabaseIdentityUninitialized
	}

	report := &dbmodel.Report{}
	if err = e.GetTx().Model(&dbmodel.Report{}).First(report, reportId).Error; err != nil {
		return nil, err
	}

	return report, nil
}

func (e *DbEnv) DBListReports(status string, accountId int64, after int64, limit int64) (*[]dbmodel.Report, error) {
	var err error

	temp := make([]dbmodel.Report, 0)
	reports := &temp

	if status != "" {
		e.SetTx(e.GetTx().Where("status = ?", status))
	}

	if accountId != 0 {
		e.SetTx(e.GetTx().Where("account_id = ?", accountId))
	}

	if after != 0 {
		e.SetTx(e.GetTx().Where("id > ?", after))
	}

	if limit != 0 {
		e.SetTx(e.GetTx().Limit(limit))
	}


	err = e.GetTx().Model(&dbmodel.Report{}).Find(reports).Error
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (e *DbEnv) DBPutReport(report *dbmodel.Report) error {
	var err error

	if report == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(&dbmodel.User{}).Create(report).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *DbEnv) DBPostReportById(report *dbmodel.Report) error {
	var err error

	if report.Id == 0 {
		return apperror.DatabaseIdentityUninitialized
	}
	if report == nil {
		return apperror.DatabaseModelUninitialized
	}

	err = e.GetTx().Model(report).Update(report).Error
	if err != nil {
		return err
	}

	return nil
}
