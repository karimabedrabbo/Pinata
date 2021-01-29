package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func PublicEntryRequestVerifyAccount(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHRequestVerifyAccount()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}

func PublicEntryConfirmVerifyAccount(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHConfirmVerifyAccount()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}


