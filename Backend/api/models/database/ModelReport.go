package dbmodel

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"time"
)

type Report struct {
	BaseModel
	HashedTokenUuid string `gorm:"size:100;not null" json:"hashed_token_uuid"`
	FromAccountId int64 `gorm:"index;not null" json:"from_account_id"`
	ToAccountId   int64 `gorm:"index;not null" json:"to_account_id"`
	ReportType   string `gorm:"size:100" json:"report_type"`
	Reason        string `gorm:"size:255" json:"reason"`
	Comment       string `gorm:"text" json:"comment"`
	ResolvedAt int64 `json:"resolved_at"`
	ReferenceBanId int64 `json:"reference_ban_id"`
	ExpiresAt int64 `gorm:"not null" json:"expires_at"`
}

func (report *Report) PrepareReportPut(r *requests.ReportPut) error {
	if r.HashedTokenUuid == "" {
		return apperror.PrepareMissingAttribute
	}
	if r.FromAccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	if r.ToAccountId == 0 {
		return apperror.PrepareMissingAttribute
	}
	if r.ReportType == "" {
		return apperror.PrepareMissingAttribute
	}
	if r.Reason == "" {
		return apperror.PrepareMissingAttribute
	}

	report.HashedTokenUuid = r.HashedTokenUuid
	report.FromAccountId = r.FromAccountId
	report.ToAccountId = r.ToAccountId
	report.ReportType = r.ReportType
	report.Reason = r.Reason
	report.Comment = r.Comment
	report.ExpiresAt = time.Now().Add(time.Hour * 24 * 60).Unix()
	report.Prepare()
	return nil
}

func (report *Report) PrepareReportPost(r *requests.ReportPost) error {
	if r.ReportId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	report.Id = r.ReportId
	report.ReferenceBanId = r.ReferenceBanId
	report.ResolvedAt = time.Now().Unix()
	report.Prepare()
	return nil
}


func (report *Report) PrepareReportPublicPost(r *requests.ReportPublicPost) error {
	if r.ReportId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	report.Id = r.ReportId
	report.ReferenceBanId = r.ReferenceBanId
	report.ResolvedAt = time.Now().Unix()
	report.Prepare()
	return nil
}