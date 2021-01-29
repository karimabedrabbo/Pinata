package endpoints

import (
	"errors"
	gjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/handlers"
	"github.com/karimabedrabbo/eyo/api/models/response"
	"time"
)

func GinJwtEntryPayloadFunc(data interface{}) gjwt.MapClaims {
	if p, ok := data.(*response.PayloadClaims); ok {
		return gjwt.MapClaims{
			k.TokenUserIdentityKey: p.UserId,
			k.TokenAccountIdentityKey: p.AccountId,
			k.TokenUniversityIdentityKey: p.UniversityId,
			k.TokenRoleKey:  p.Role,
			k.TokenEmailVerifiedKey:  p.EmailVerified,
			k.TokenUniversityVerified: p.UniversityVerified,
		}
	}

	return gjwt.MapClaims{}
}

func GinJwtEntryAuthenticator(c *gin.Context) (interface{}, error) {
	var err error
	h := handlers.GetRequestHandler(c)

	var payloadClaims *response.PayloadClaims
	switch h.C.MustGet(k.GinJwtLoginTypeKey) {
	case k.GinJwtLoginTypeIdentity:
		if payloadClaims, err = h.RHGinJwtIdentityAuthenticator(); err != nil {
			return nil, err
		}
	case k.GinJwtLoginTypeNoIdentity:
		if payloadClaims, err = h.RHGinJwtAnonymousAuthenticator(); err != nil {
			return nil, err
		}
	}

	return payloadClaims, nil
}

func GinJwtEntryUnauthorized(c *gin.Context, code int, message string) {
	c.Set(k.GinJwtStatusCode, code)
	c.Set(k.GinJwtErrorMessage, message)

	//we need to do this (mandatory) because this is called on all token checks, not just login handlers
	//if we don't abort here we actually allow invalid/no-show tokens through
	handlers.GetRequestHandler(c).RHAbortRequest(code, errors.New(message))
}

func GinJwtEntryLoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.Set(k.GinJwtStatusCode, code)
	c.Set(k.GinJwtToken, token)
	c.Set(k.GinJwtExpireAt, expire.Nanosecond())
}

func GinJwtEntryLogoutResponse(c *gin.Context, code int) {
	c.Set(k.GinJwtStatusCode, code)
}

func GinJwtEntryRefreshResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.Set(k.GinJwtStatusCode, code)
	c.Set(k.GinJwtToken, token)
	c.Set(k.GinJwtExpireAt, expire.Nanosecond())
}
