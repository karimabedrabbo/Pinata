package managers

import (
	"github.com/gin-gonic/gin"
	"github.com/karimabedrabbo/eyo/api/apputils"
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/endpoints"
)

var router *Router

type Router struct {
	Engine *gin.Engine
}

func SetupRouter() *Router {
	return &Router{
		Engine: setupEngine(),
	}
}

func InitRouter() {
	router = SetupRouter()
}

func GetRouter() *Router {
	return router
}

func setupEngine() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 10 * 1024 * 1024
	router.Use(HandlerEnviornmentMiddleware())
	router.Use(CORSMiddleware())

	//not currently working within docker
	//router.ForwardedByClientIP = true
	//if apputils.GetAppEnvIsProduction() {
	//	router.Use(FilterIpMiddleware())
	//}

	setupRoutes(router)
	return router
}

func setupRoutes(e *gin.Engine)  {
	api := e.Group(k.AppApiPath)
	setupRoutesGroupPublic(api)
	setupRoutesGroupAnonymous(api)
	setupRoutesGroupUser(api)
	setupRoutesGroupSysadmin(api)
}

func setupRoutesGroupPublic(api *gin.RouterGroup) {
	public := api.Group(k.RouteDirPublic)
	if apputils.GetAppEnvIsProduction() {
		public.Use(PublicRateLimiterMiddleware())
	}

	//review privileged operations
	//todo
	public.POST(k.RouteReport + k.RouteSpecifyReport, endpoints.PublicReviewReport)

	//verify email
	public.PUT(k.RouteVerifyAccount, endpoints.PublicEntryRequestVerifyAccount)
	public.POST(k.RouteVerifyAccount, endpoints.PublicEntryConfirmVerifyAccount)

	//forgot password
	public.PUT(k.RoutePasswordReset, endpoints.PublicEntryRequestPasswordReset)
	public.POST(k.RoutePasswordReset, endpoints.PublicEntryConfirmPasswordReset)

	//login and signup
	public.POST(k.RouteLogin, endpoints.PublicEntryLoginUser)
	public.POST(k.RouteLoginAnonymous, endpoints.PublicEntryLoginAnonymousUser)
	public.POST(k.RouteSignup, endpoints.PublicEntrySignupUser)

	public.POST(k.RouteAttachment, endpoints.PublicEntryConfirmAttachment)
}

func setupRoutesGroupAnonymous(api *gin.RouterGroup) {
	anonymous := api.Group(k.RouteDirAnonymous)
	anonymous.Use(AuthenticationMiddleware())
	anonymous.Use(AnonymousAuthorizeMiddleware())
	if apputils.GetAppEnvIsProduction() {
		anonymous.Use(IdentityRateLimiterMiddleware())
	}

	//university public
	anonymous.GET(k.RouteUniversity + k.RouteSpecifyUniversity, endpoints.AnonymousEntryGetUniversity)
	anonymous.GET(k.RouteUniversity + k.RouteActionList, endpoints.AnonymousEntryGetUniversityList)
	anonymous.GET(k.RouteUniversity + k.RouteSpecifyUniversity + k.RouteAvatar, endpoints.AnonymousEntryGetUniversityAvatar)

	//pairing anonymous
	//anonymous.GET(k.RouteUniversity + k.RouteSpecifyUniversity + k.RoutePairing, endpoints.PublicEntryGetPairingBoard)

}
func setupRoutesGroupUser(api *gin.RouterGroup) {
	user := api.Group(k.RouteDirUser)
	user.Use(AuthenticationMiddleware())
	user.Use(UserAuthorizeMiddleware())
	if apputils.GetAppEnvIsProduction() {
		user.Use(IdentityRateLimiterMiddleware())
	}

	//logistical things
	//todo
	user.PUT(k.RouteReport, endpoints.UserEntrySendReport)
	user.PUT(k.RouteFeedback, endpoints.UserEntrySendFeedback)
	user.POST(k.RouteChangePassword, endpoints.UserEntryChangePassword)
	user.POST(k.RouteChangeEmail, endpoints.UserEntryChangeEmail)
	user.POST(k.RouteCloseAccount, endpoints.UserEntryCloseAccount)
	user.POST(k.RouteLogout, endpoints.UserEntryLogoutUser)

	//user profile
	user.GET(k.RouteUser + k.RouteProfile, endpoints.UserEntryGetUserProfile)
	user.GET(k.RouteUser + k.RouteProfile + k.RouteAvatar, endpoints.UserEntryGetUserAvatar)
	user.POST(k.RouteUser + k.RouteProfile, endpoints.UserEntryUpdateUserProfile)
	user.POST(k.RouteUser + k.RouteProfile + k.RouteAvatar, endpoints.UserEntryUpdateUserAvatar)


	//user attachments
	//user.GET(k.RouteUser + k.RouteAttachment + k.RouteActionList, endpoints.UserEntryGetUserAttachmentList)
	//user.GET(k.RouteUser + k.RouteAttachment + k.RouteSpecifyAttachment, endpoints.UserEntryGetUserAttachment)
	//user.GET(k.RouteUser + k.RouteAvatar, endpoints.UserEntryGetUserAvatar)
	//user.PUT(k.RouteUser + k.RouteAvatar, endpoints.UserEntryCreateUserAvatar)


	//user blocking
	//user.PUT(k.RouteBlock, endpoints.UserEntryBlockUser)
	//user.GET(k.RouteBlock + k.RouteSpecifyUser, endpoints.UserEntryGetBlockUser)
	//user.GET(k.RouteBlock + k.RouteActionList, endpoints.UserEntryGetBlockListUser)
	//user.DELETE(k.RouteBlock + k.RouteSpecifyUser, endpoints.UserEntryUnblockUser)


	//university
	user.GET(k.RouteUniversity + k.RouteSpecifyUniversity, endpoints.UserEntryGetUniversity)
	user.GET(k.RouteUniversity + k.RouteActionList, endpoints.UserEntryGetUniversityList)
	user.GET(k.RouteUniversity + k.RouteSpecifyUniversity + k.RouteAvatar, endpoints.UserEntryGetUniversityAvatar)

	//pairing
	//user.GET(k.RouteUniversity + k.RouteSpecifyUniversity + k.RoutePairing, endpoints.UserEntryGetPairingBoard)
	//user.PUT(k.RoutePairing, endpoints.UserEntryRequestPairing)
	//user.POST(k.RoutePairing + k.RouteSpecifyPairing, endpoints.UserEntryUpdatePairing)
	//user.DELETE(k.RoutePairing + k.RouteSpecifyPairing, endpoints.UserEntryDeletePairing)


	//conversations
	//user.PUT(k.RouteConversation, endpoints.UserEntryCreateConversation)
	//user.POST(k.RouteConversation + k.RouteSpecifyConversation, endpoints.UserEntryUpdateConversation)
	//user.DELETE(k.RouteConversation + k.RouteSpecifyConversation, endpoints.UserEntryLeaveConversation)
	//user.GET(k.RouteConversation + k.RouteSpecifyConversation, endpoints.UserEntryGetConversation)
	//user.GET(k.RouteConversation + k.RouteActionList, endpoints.UserEntryGetConversationList)


	//conversation attachments
	//user.GET(k.RouteConversation + k.RouteSpecifyConversation + k.RouteAttachment + k.RouteActionList, endpoints.UserEntryGetConversationAttachmentList)
	//user.GET(k.RouteConversation + k.RouteSpecifyConversation + k.RouteAttachment + k.RouteSpecifyConversationAttachment, endpoints.UserEntryGetConversationAttachment)
	//user.GET(k.RouteConversation + k.RouteSpecifyConversation + k.RouteAvatar, endpoints.UserEntryGetConversationAvatar)
	//user.PUT(k.RouteConversation + k.RouteSpecifyConversation + k.RouteAvatar, endpoints.UserEntryCreateConversationAvatar)


	//conversation participants
	//user.GET(k.RouteConversation + k.RouteSpecifyConversation + k.RouteParticipant + k.RouteActionList, endpoints.UserEntryGetConversationParticipantList)
	//user.PUT(k.RouteConversation + k.RouteSpecifyConversation + k.RouteParticipant, endpoints.UserEntryAddConversationParticipant)
	//user.POST(k.RouteConversation + k.RouteSpecifyConversation + k.RouteParticipant + k.RouteSpecifyConversationParticipant, endpoints.UserEntryUpdateConversationParticipant)
	//user.GET(k.RouteConversation + k.RouteSpecifyConversation + k.RouteParticipant + k.RouteSpecifyConversationParticipant, endpoints.UserEntryGetConversationParticipant)
	//user.DELETE(k.RouteConversation + k.RouteSpecifyConversation + k.RouteParticipant + k.RouteSpecifyConversationParticipant, endpoints.UserEntryRemoveConversationParticipant)
	//user.GET(k.RouteUser + k.RouteSpecifyConversation + k.RouteParticipant + k.RouteProfile, endpoints.UserEntryGetConversationParticipantProfile)
	//user.GET(k.RouteUser + k.RouteSpecifyConversation + k.RouteParticipant + k.RouteAvatar, endpoints.UserEntryGetConversationParticipantAvatar)


	//conversation messages
	//user.GET(k.RouteConversation + k.RouteSpecifyConversation + k.RouteMessage + k.RouteActionList, endpoints.UserEntryGetConversationMessageList)
	//user.PUT(k.RouteConversation + k.RouteSpecifyConversation + k.RouteMessage, endpoints.UserEntryCreateConversationMessage)
	//user.GET(k.RouteConversation + k.RouteSpecifyConversation + k.RouteMessage + k.RouteSpecifyConversationMessage, endpoints.UserEntryGetConversationMessage)

}

func setupRoutesGroupSysadmin(api *gin.RouterGroup) {
	sysadmin := api.Group(k.RouteDirSysadmin)
	sysadmin.Use(AuthenticationMiddleware())
	sysadmin.Use(SysadminAuthorizeMiddleware())
	if apputils.GetAppEnvIsProduction() {
		sysadmin.Use(IdentityRateLimiterMiddleware())
	}

	//attachment privileged operation
	sysadmin.GET(k.RouteAttachment + k.RouteSpecifyAttachment, endpoints.SysadminEntryGetAttachment)
	sysadmin.GET(k.RouteAttachment + k.RouteActionList, endpoints.SysadminEntryGetAttachmentList)
	sysadmin.DELETE(k.RouteAttachment + k.RouteSpecifyAttachment, endpoints.SysadminEntryDeleteAttachment)

	//university privileged operations
	sysadmin.PUT(k.RouteUniversity, endpoints.SysadminEntryPutUniversity)
	sysadmin.POST(k.RouteUniversity + k.RouteSpecifyUniversity + k.RouteAvatar, endpoints.SysadminEntryPostUniversityAvatar)
	sysadmin.POST(k.RouteUniversity + k.RouteSpecifyUniversity, endpoints.SysadminEntryPostUniversity)
	sysadmin.DELETE(k.RouteUniversity + k.RouteSpecifyUniversity, endpoints.SysadminEntryDeleteUniversity)


	//ban privileged operations
	//todo
	sysadmin.PUT(k.RouteBan, endpoints.SysadminEntryBanUser)
	sysadmin.GET(k.RouteBan + k.RouteSpecifyUser, endpoints.SysadminEntryGetBan)
	sysadmin.DELETE(k.RouteBan + k.RouteSpecifyUser, endpoints.SysadminEntryUnbanUser)

	//pairing privileged operations
	//sysadmin.PUT(k.RoutePairing + k.RouteSpecifyPairing + k.RouteAttribute, endpoints.SysadminEntryPutPairingAttribute)
	//sysadmin.POST(k.RoutePairing + k.RouteSpecifyPairing + k.RouteAttribute + k.RouteSpecifyPairingAttribute, endpoints.SysadminEntryPostPairingAttribute)
	//sysadmin.DELETE(k.RoutePairing + k.RouteSpecifyPairing + k.RouteAttribute + k.RouteSpecifyPairingAttribute, endpoints.SysadminEntryDeleteAttribute)


	//report privileged operations
	//todo
	sysadmin.GET(k.RouteReport + k.RouteSpecifyReport, endpoints.EntrySysadminGetReport)
	sysadmin.GET(k.RouteReport + k.RouteActionList, endpoints.EntrySysadminGetReportList)
	sysadmin.POST(k.RouteReport + k.RouteSpecifyReport, endpoints.EntrySysadminUpdateReportStatus)

}


