package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func UserEntrySendFeedback(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHPutFeedback()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}


