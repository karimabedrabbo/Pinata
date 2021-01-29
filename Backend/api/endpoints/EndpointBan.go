package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func SysadminEntryGetBan(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, ban, err := h.RHGetBan()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, ban)
}

func SysadminEntryBanUser(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHPutBan()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}


func SysadminEntryUnbanUser(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHDeleteBan()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}