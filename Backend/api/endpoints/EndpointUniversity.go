package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/handlers"
)

func AnonymousEntryGetUniversity(c *gin.Context)  {
	h := handlers.GetRequestHandler(c)
	code, univ, err := h.RHGetUniversity()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, univ)
}


func AnonymousEntryGetUniversityAvatar(c *gin.Context)  {
	h := handlers.GetRequestHandler(c)
	code, url, err := h.RHGetUniversityAvatar()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, url)
}

func AnonymousEntryGetUniversityList(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, univs, err := h.RHGetUniversityList()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, univs)
}

func UserEntryGetUniversity(c *gin.Context)  {
	h := handlers.GetRequestHandler(c)
	code, univ, err := h.RHGetUniversity()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, univ)
}


func UserEntryGetUniversityAvatar(c *gin.Context)  {
	h := handlers.GetRequestHandler(c)
	code, url, err := h.RHGetUniversityAvatar()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, url)
}

func UserEntryGetUniversityList(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, univs, err := h.RHGetUniversityList()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, univs)
}


func SysadminEntryPutUniversity(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHPutUniversity()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}

func SysadminEntryPostUniversityAvatar(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, url, err := h.RHPostUniversityAvatar()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, url)
}

func SysadminEntryPostUniversity(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHPostUniversity()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}

func SysadminEntryDeleteUniversity(c *gin.Context) {
	h := handlers.GetRequestHandler(c)
	code, err := h.RHDeleteUniversity()
	if err != nil {
		h.RHAbortRequest(code, err)
		return
	}

	h.RHOkRequest(code, nil)
}
