package dbmodel

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/models/requests"
)

type Ban struct {
	BaseModel
	IssuerAccountId int64 `gorm:"index;not null" json:"issuer_account_id"`
	ToAccountId int64 `gorm:"index;not null" json:"to_account_id"`
}


func (ban *Ban) PrepareReportPost(r *requests.ReportPost) error {
	if r.IssuerAccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	if r.BanAccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	ban.IssuerAccountId = r.IssuerAccountId
	ban.ToAccountId = r.BanAccountId
	ban.Prepare()
	return nil
}

func (ban *Ban) PrepareReportPublicPost(r *requests.ReportPublicPost) error {
	if r.IssuerAccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	if r.BanAccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	ban.IssuerAccountId = r.IssuerAccountId
	ban.ToAccountId = r.BanAccountId
	ban.Prepare()
	return nil
}


func (ban *Ban) PrepareBanPut(r *requests.BanPut) error {
	if r.IssuerAccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	if r.ToAccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	ban.IssuerAccountId = r.IssuerAccountId
	ban.ToAccountId = r.ToAccountId
	ban.Prepare()
	return nil
}