package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func EntrySysadminGetReport(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, report, err := h.RHGetReport()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, report)
}

func EntrySysadminGetReportList(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, reports, err := h.RHGetReportList()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, reports)
}


func UserEntrySendReport(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHPutReport()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}

func PublicReviewReport(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHPostReport()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}


func EntrySysadminUpdateReportStatus(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHPostReportPublic()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}
