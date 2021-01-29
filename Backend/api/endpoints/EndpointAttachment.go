package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func PublicEntryConfirmAttachment(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHConfirmAttachment()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}


func SysadminEntryGetAttachment(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, url, err := h.RHGetAttachment()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, url)
}

func SysadminEntryGetAttachmentList(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, urls, err := h.RHGetAttachmentList()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, urls)
}

func SysadminEntryDeleteAttachment(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, url, err := h.RHDeleteAttachment()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, url)
}
