package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func UserEntryGetUserProfile(c *gin.Context)  {
	h := handlers.GetRequestHandler(c)
	code, user, err := h.RHGetUserProfile()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, user)
}

func UserEntryGetUserAvatar(c *gin.Context)  {
	h := handlers.GetRequestHandler(c)
	code, url, err := h.RHGetUserAvatar()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, url)
}


func UserEntryUpdateUserProfile(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHPostUserProfile()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}


func UserEntryUpdateUserAvatar(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, url, err := h.RHPostUserAvatar()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, url)
}