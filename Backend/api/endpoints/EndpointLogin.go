package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func PublicEntryLoginUser(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, payloadToken, err := h.RHLoginUser()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, payloadToken)
}

func PublicEntryLoginAnonymousUser(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, payloadToken, err := h.RHLoginAnonymousUser()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, payloadToken)
}

func UserEntryLogoutUser(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHLogoutUser()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}
