package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func PublicEntryRequestPasswordReset(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHRequestPasswordReset()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}

func PublicEntryConfirmPasswordReset(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHConfirmPasswordReset()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}

