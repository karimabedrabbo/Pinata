package handlers

import (
	"github.com/google/uuid"
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/apputils"
	k "github.com/karimabedrabbo/eyo/api/constants"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"github.com/karimabedrabbo/eyo/api/models/requests"
	"net/http"
	"time"
)

func (h *RhEnv) RHGetReport() (int, *dbmodel.Report, error) {
	var err error

	reportGet := &requests.ReportGet{}

	if err = h.MapUri(reportGet); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var report *dbmodel.Report
	if report, err = h.E.DBGetReportById(reportGet.ReportId); err != nil {
		return http.StatusNotFound, nil, err
	}

	return http.StatusOK, report, nil
}


func (h *RhEnv) RHGetReportList() (int, *[]dbmodel.Report, error) {
	var err error

	reportList := &requests.ReportList{}
	if err = h.MapQuery(reportList); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var reports *[]dbmodel.Report
	if reports, err = h.E.DBListReports(reportList.Status, reportList.AccountId, reportList.AfterId, reportList.Limit); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, reports, nil
}

func (h *RhEnv) RHPutReport() (int, error) {
	var err error

	reportPut := &requests.ReportPut{}
	if err = h.MapJson(reportPut); err != nil {
		return http.StatusBadRequest, err
	}

	var fromAccountId int64
	if fromAccountId = h.Auth.GetAccountId(h.C); fromAccountId == 0 {
		return http.StatusBadRequest, apperror.RequestTokenIdentityMissing
	}

	var tokenUuid uuid.UUID
	var tokenHashedUuid string
	if tokenUuid, err = apputils.GetUUID(); err != nil {
		return http.StatusInternalServerError, err
	}
	if tokenHashedUuid, err = apputils.DerivePasswordHash(tokenUuid.String()); err != nil {
		return http.StatusInternalServerError, err
	}

	reportPut.FromAccountId = fromAccountId
	reportPut.HashedTokenUuid = tokenHashedUuid

	report := &dbmodel.Report{}
	if err = report.PrepareReportPut(reportPut); err != nil {
		return http.StatusBadRequest, err
	}

	if err = h.E.DBPutReport(report); err != nil {
		return http.StatusInternalServerError, err
	}


	baseReportUrl := apputils.GetAppUrl() + k.RouteVerifyAccount + "?report_id=" + apputils.IdToString(report.Id) + "&token_uuid=" + tokenUuid.String()
	noActionResolveUrl := baseReportUrl + "&decision=no_action"
	banOffenderUrl := baseReportUrl + "&decision=ban_offender"
	banSenderUrl := baseReportUrl + "&decision=ban_sender"

	//todo fill in message history
	if err = h.Mail.SendReport(report.Reason, `{"history": "test message history"}`, noActionResolveUrl, banOffenderUrl, banSenderUrl); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil

}

func (h *RhEnv) RHPostReport() (int, error) {
	var err error

	reportPost := &requests.ReportPost{}
	if err = h.MapUriJson(reportPost); err != nil {
		return http.StatusBadRequest, err
	}

	if reportPost.IssuerAccountId = h.Auth.GetAccountId(h.C); reportPost.IssuerAccountId != 0 {
		return http.StatusBadRequest, apperror.RequestTokenIdentityMissing
	}

	var report *dbmodel.Report
	if report, err = h.E.DBGetReportById(reportPost.ReportId); err != nil {
		return http.StatusNotFound, err
	}

	if reportPost.Decision != k.ReportDecisionNoAction {
		switch reportPost.Decision {
		case k.ReportDecisionBanOffender:
			reportPost.BanAccountId = report.ToAccountId
		case k.ReportDecisionBanSender:
			reportPost.BanAccountId = report.FromAccountId
		default:
			return http.StatusBadRequest, apperror.RequestNotSupported
		}

		ban := &dbmodel.Ban{}
		if err = ban.PrepareReportPost(reportPost); err != nil {
			return http.StatusBadRequest, err
		}

		if err = h.E.DBPutBan(ban); err != nil {
			return http.StatusInternalServerError, err
		}

		//assign the newly created ban
		reportPost.ReferenceBanId = ban.Id
	}

	if err = report.PrepareReportPost(reportPost); err != nil {
		return http.StatusInternalServerError, err
	}

	if err = h.E.DBPostReportById(report); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (h *RhEnv) RHPostReportPublic() (int, error) {
	var err error

	reportPublicPost := &requests.ReportPublicPost{}
	if err = h.MapUriJson(reportPublicPost); err != nil {
		return http.StatusBadRequest, err
	}

	if reportPublicPost.IssuerAccountId = h.Auth.GetAccountId(h.C); reportPublicPost.IssuerAccountId != 0 {
		return http.StatusBadRequest, apperror.RequestTokenIdentityMissing
	}

	var report *dbmodel.Report
	if report, err = h.E.DBGetReportById(reportPublicPost.ReportId); err != nil {
		return http.StatusNotFound, err
	}

	if err = apputils.VerifyPassword(report.HashedTokenUuid, reportPublicPost.TokenUuid.String()); err != nil {
		return http.StatusUnauthorized, apperror.RequestInvalidCredentials
	}

	if report.ExpiresAt > time.Now().Unix() {
		return http.StatusUnauthorized, apperror.RequestWindowExpired
	}

	if report.ResolvedAt > 0 {
		return http.StatusUnauthorized, apperror.RequestAlreadyUsed
	}

	if reportPublicPost.Decision != k.ReportDecisionNoAction {
		switch reportPublicPost.Decision {
		case k.ReportDecisionBanOffender:
			reportPublicPost.BanAccountId = report.ToAccountId
		case k.ReportDecisionBanSender:
			reportPublicPost.BanAccountId = report.FromAccountId
		default:
			return http.StatusBadRequest, apperror.RequestNotSupported
		}

		ban := &dbmodel.Ban{}
		if err = ban.PrepareReportPublicPost(reportPublicPost); err != nil {
			return http.StatusBadRequest, err
		}

		if err = h.E.DBPutBan(ban); err != nil {
			return http.StatusInternalServerError, err
		}

		//assign the newly created ban
		reportPublicPost.ReferenceBanId = ban.Id
	}

	if err = report.PrepareReportPublicPost(reportPublicPost); err != nil {
		return http.StatusInternalServerError, err
	}

	if err = h.E.DBPostReportById(report); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}