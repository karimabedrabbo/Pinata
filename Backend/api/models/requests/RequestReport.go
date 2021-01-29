package requests

import "github.com/google/uuid"

type ReportGet struct {
	ReportId int64 `uri:"report_id" json:"report_id" binding:"required,numeric"`
}

type ReportList struct {
	ListRequest
	AccountId int64 `json:"account_id" binding:"omitempty,numeric"`
	Status string `json:"status" binding:"omitempty,oneof=queued processing positive negative"`
}

type ReportPut struct {
	HashedTokenUuid string `json:"-" binding:"-"`
	FromAccountId int64 `json:"-" binding:"-"`
	ToAccountId int64 `json:"to_account_id" binding:"required,numeric"`
	ReportType string `json:"report_type" binding:"required,oneof=conversation account"`
	Reason string `json:"reason" binding:"required,oneof=dangerous_illegal spam_scam_fake obscene_sexual hate_speech_personal_attack"`
	Comment string `json:"comment" binding:"omitempty,max=3000"`
}

type ReportPost struct {
	ReportId int64 `uri:"report_id" json:"report_id" binding:"required,numeric"`
	Decision string `json:"status" binding:"required,oneof=ban_offender ban_sender no_action"`
	ReferenceBanId int64 `json:"-" binding:"-"`
	IssuerAccountId int64 `json:"-" binding:"-"`
	BanAccountId int64 `json:"-" binding:"-"`
}

type ReportPublicPost struct {
	ReportId int64 `form:"report_id" json:"report_id" binding:"required,numeric"`
	TokenUuid uuid.UUID `form:"token_uuid" json:"token_uuid" binding:"required,uuid"`
	Decision string `json:"status" binding:"required,oneof=ban_offender ban_sender no_action"`
	ReferenceBanId int64 `json:"-" binding:"-"`
	IssuerAccountId int64 `json:"-" binding:"-"`
	BanAccountId int64 `json:"-" binding:"-"`
}



