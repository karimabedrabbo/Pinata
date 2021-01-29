package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/database"
	"github.com/karimabedrabbo/eyo/api/models/response"
	"net/http"
)

//request handler enviornment
type RhEnv struct {
	E *database.DbEnv
	C *gin.Context
	Mail Mails
	Sanitize Sanitizes
	Storage Stores
	Auth Authenticates
}

type Mails interface {
	SendReport(reason string, messageHistoryPrettyJson string, noActionResolveUrl string, banOffenderUrl string, banSenderUrl string) error
	SendFeedback(name string, email string) error
	SendVerifyEmail(name string, email string, verifyUrl string) error
	SendPasswordReset(name string, email string, resetUrl string) error
}

type Sanitizes interface {
	SanitizeJsonBytes(b []byte) (interface{}, error)
	SanitizeJsonObj(obj interface{}) (interface{}, error)
	SanitizeString(str string) string
}

type Stores interface {
	GetSignedUrl(method string, mediaType string, encodingType string, md5 string, attachmentId int64, tokenUuid uuid.UUID) (*response.SignedUrlResponse, error)
}

type Authenticates interface {
	GetClaims(c *gin.Context) map[string]interface{}
	GetRole(c *gin.Context) string
	GetUniversityId(c *gin.Context) int64
	GetUserId(c *gin.Context) int64
	GetAccountId(c *gin.Context) int64
	LoginHandler(c *gin.Context)
	LogoutHandler(c *gin.Context)
}

func (h *RhEnv) internalMapUri(r interface{}) error{
	return h.C.ShouldBindUri(r)
}

func (h *RhEnv) internalMapJson(r interface{}) error {
	return h.C.ShouldBindWith(r, binding.JSON)
}

func (h *RhEnv) internalMapQuery(r interface{}) error {
	return h.C.ShouldBindWith(r, binding.Query)
}

func (h *RhEnv) MapUri(r interface{}) error {
	var err error
	if err = h.internalMapUri(r); err != nil {
		return err
	}
	return nil
}

func (h *RhEnv) MapJson(r interface{}) error {
	var err error
	if err = h.internalMapJson(r); err != nil {
		return err
	}
	return nil
}

func (h *RhEnv) MapQuery(r interface{}) error {
	var err error
	if err = h.internalMapQuery(r); err != nil {
		return err
	}
	return nil
}

func (h *RhEnv) MapUriQuery(r interface{}) error {
	var err error
	if err = h.internalMapQuery(r); err != nil {
		return err
	}
	if err = h.internalMapUri(r); err != nil {
		return err
	}
	return nil
}

func (h *RhEnv) MapUriJson(r interface{}) error {
	var err error
	if err = h.internalMapJson(r); err != nil {
		return err
	}
	if err = h.internalMapUri(r); err != nil {
		return err
	}
	return nil
}

func (h *RhEnv) MapUriQueryJson(r interface{}) error {
	var err error
	if err = h.internalMapQuery(r); err != nil {
		return err
	}
	if err = h.internalMapJson(r); err != nil {
		return err
	}
	if err = h.internalMapUri(r); err != nil {
		return err
	}
	return nil
}

func (h *RhEnv) RHAbortRequest(code int, err error) {
	//rollback done in handler middleware (which allows for other context aborts to still rollback the transaction)
	//h.E.GetTx().Rollback()
	if !h.C.IsAborted() {

		h.C.AbortWithStatusJSON(code, gin.H{
			"code":  code,
			"error": h.Sanitize.SanitizeString(err.Error()),
		})
	}

}

func (h *RhEnv) RHOkRequest(code int, obj interface{}) {
	if !h.C.IsAborted() {
		h.E.GetTx().Commit()
		var err error
		if obj, err = h.Sanitize.SanitizeJsonObj(obj); err != nil {
			h.RHAbortRequest(http.StatusInternalServerError, err)
			return
		}
		h.C.JSON(code, obj)
	}
}

func (h *RhEnv) FinalizeRequestHandler() {
	h.E.GetTx().RollbackUnlessCommitted()
}

func GetRequestHandler(c *gin.Context) *RhEnv  {
	return c.MustGet(k.RequestHandlerKey).(*RhEnv)
}


