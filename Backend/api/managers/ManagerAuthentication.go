package managers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/endpoints"
	"log"
	"time"
)

type Auth struct {
	GinJwtClient *jwt.GinJWTMiddleware
}

var auth *Auth

func SetupAuthentication() *Auth {

	tempAuth, err := jwt.New(&jwt.GinJWTMiddleware{
		Timeout:time.Hour * 24,
		Realm: k.GinJwtAuthenticationRealm,
		Key:         apputils.GetAppApiSecret(),
		Authenticator: endpoints.GinJwtEntryAuthenticator,
		PayloadFunc: endpoints.GinJwtEntryPayloadFunc,
		LoginResponse: endpoints.GinJwtEntryLoginResponse,
		LogoutResponse: endpoints.GinJwtEntryLogoutResponse,
		RefreshResponse: endpoints.GinJwtEntryRefreshResponse,
		Unauthorized: endpoints.GinJwtEntryUnauthorized,
		IdentityKey: k.TokenUserIdentityKey,
		DisabledAbort: true, //we roll our own aborting
	})
	if err != nil {
		log.Fatalf("error setting up authentication: %v", err)
	}

	return &Auth {
		GinJwtClient: tempAuth,
	}
}

func InitAuthentication() {
	auth = SetupAuthentication()
}

func GetAuthentication() *Auth {
	return auth
}

func (a *Auth) GetClaims(c *gin.Context) map[string]interface{} {
	return jwt.ExtractClaims(c)
}

func (a *Auth) GetRole(c *gin.Context) string {
	claims := a.GetClaims(c)
	if role, exists := claims[k.TokenRoleKey].(string); exists {
		return role
	}
	return ""
}

func (a *Auth) GetUniversityId(c *gin.Context) int64 {
	claims := a.GetClaims(c)
	if role, exists := claims[k.TokenUniversityIdentityKey].(int64); exists {
		return role
	}
	return 0
}

func (a *Auth) GetUserId(c *gin.Context) int64 {
	claims := a.GetClaims(c)
	if role, exists := claims[k.TokenUserIdentityKey].(int64); exists {
		return role
	}
	return 0
}

func (a *Auth) GetAccountId(c *gin.Context) int64 {
	claims := a.GetClaims(c)
	if role, exists := claims[k.TokenAccountIdentityKey].(int64); exists {
		return role
	}
	return 0
}

func (a *Auth) LoginHandler(c *gin.Context) {
	auth.GinJwtClient.LoginHandler(c)
}

func (a *Auth) LogoutHandler(c *gin.Context) {
	auth.GinJwtClient.LogoutHandler(c)
}