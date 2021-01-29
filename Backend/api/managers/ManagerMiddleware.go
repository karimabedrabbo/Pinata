package managers

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/handlers"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"net/http"
)

func HandlerEnviornmentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		InitRequestHandler(c)
		c.Next()
		handlers.GetRequestHandler(c).FinalizeRequestHandler()
	}
}


func AuthenticationMiddleware() gin.HandlerFunc {
	return GetAuthentication().GinJwtClient.MiddlewareFunc()
}

func AnonymousAuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if GetAuthentication().GetRole(c) != k.AccountRoleAnonymous {
			handlers.GetRequestHandler(c).RHAbortRequest(http.StatusUnauthorized, apperror.RequestTokenIdentityInvalid)
		}
		c.Next()
	}
}

func UserAuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if GetAuthentication().GetRole(c) != k.AccountRoleUser {
			handlers.GetRequestHandler(c).RHAbortRequest(http.StatusUnauthorized, apperror.RequestTokenIdentityInvalid)
		}
		c.Next()
	}
}

func SysadminAuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if GetAuthentication().GetRole(c) != k.AccountRoleSysadmin {
			handlers.GetRequestHandler(c).RHAbortRequest(http.StatusUnauthorized, apperror.RequestTokenIdentityInvalid)
		}
		c.Next()
	}
}


// This enables us interact with the React Frontend
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			handlers.GetRequestHandler(c).RHAbortRequest(http.StatusNoContent, apperror.RequestNotSupported)
			return
		}
		c.Next()
	}
}

func PublicRateLimiterMiddleware() gin.HandlerFunc {
	return mgin.NewMiddleware(GetRateLimit().publicLimiterClient)
}


func IdentityRateLimiterMiddleware() gin.HandlerFunc {
	return mgin.NewMiddleware(GetRateLimit().identityLimiterClient)
}

func FilterIpMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !GetIpFilter().IpAllowed(c.ClientIP()) {
			handlers.GetRequestHandler(c).RHAbortRequest(http.StatusForbidden, apperror.RequestForbidden)
			return
		}

		c.Next()
	}
}


//func RequestSizeLimiter() gin.HandlerFunc {
//	var maxFileSize int64 = 10 * 1024 * 1024
//	return limits.RequestSizeLimiter(maxFileSize)
//}
