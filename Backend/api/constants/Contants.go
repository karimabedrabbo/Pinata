package k


//application constants
var (
	AppName = "Eyo"
	AppEnvProduction = "production"
	AppEnvDevelopment = "development"
	AppEnvKeyPrefixProduction = "PROD_"
	AppEnvKeyPrefixDevelopment = "DEV_"
	AppApiPath = "/api/v1"
	AppSupportEmail = "support@geteyo.com"
	AppProductionUrl = "https://geteyo.com"
	AppDevelopmentUrl = "http://127.0.0.1"
	AppContextRemote = "remote"
   	AppContextLocal = "local"
	AppContextLocalContainer = "local_container"
	AppContextKeyPrefixLocal = "LOCAL_"
	AppContextKeyPrefixLocalContainer = "LOCAL_CONTAINER_"
	AppContextKeyPrefixRemote = "REMOTE_"
)


//database
var (
	DatabaseKeyInitialPrefix = "DB_"
	DatabaseKeySuffixHost = "HOST"
	DatabaseKeySuffixUser = "USER"
	DatabaseKeySuffixPort = "PORT"
	DatabaseKeySuffixName = "NAME"
	DatabaseKeySuffixPassword = "PASSWORD"
	DatabaseKeySuffixEnv = "ENV"

)

//request handling constants
var (
	RequestHandlerKey = "request_handler"
)

//jwt token
var (
	TokenUserIdentityKey = "user_id"
	TokenAccountIdentityKey = "account_id"
	TokenRoleKey = "role"
	TokenEmailVerifiedKey = "email_verified"
	TokenUniversityIdentityKey = "university_id"
	TokenUniversityVerified = "university_verified"

)

//authentication manager
var (
	GinJwtAuthenticationRealm = "geteyo.com"
	GinJwtLoginTypeKey = "gin_jwt_login_type"
	GinJwtLoginTypeNoIdentity = "no_identity"
	GinJwtLoginTypeIdentity = "identity"
	GinJwtStatusCode = "gin_jwt_status_code"
	GinJwtToken = "gin_jwt_token"
	GinJwtExpireAt = "gin_jwt_expire_at"
	GinJwtErrorMessage = "gin_jwt_error_message"
)


//rate limiter
var (
	RedisKeyInitialPrefix = "REDIS_"
	RedisKeySuffixHost = "HOST"
	RedisKeySuffixPort = "PORT"
	RedisRateLimiterPublicPrefix = "public_rate_limiter"
	RedisRateLimiterIdentityPrefix = "identity_rate_limiter"
)

var (
	AccountRoleAnonymous = "anonymous"
	AccountRoleUser = "user"
	AccountRoleSysadmin = "sysadmin"
)


//storage manager
var (
	StorageBucketName = "eyobucket"
)


//Routes
var (
	RouteDirPublic = "/public"
	RouteDirAnonymous = "/anonymous"
	RouteDirUser = "/user"
	RouteDirSysadmin = "/sysadmin"
	RouteVerifyAccount = "/verify_account"
	RoutePasswordReset = "/password_reset"
	RouteLogin = "/login"
	RouteLoginAnonymous = "/login_anonymous"
	RouteSignup = "/signup"
	RouteLogout = "/logout"
	RouteUniversity = "/university"
	RouteChangePassword = "/change_password"
	RouteChangeEmail = "/change_email"
	RouteCloseAccount = "/close_account"
	RouteUser = "/user"
	RouteProfile = "/profile"
	RouteAvatar = "/avatar"
	RouteAttachment = "/attachment"
	RouteConversation = "/conversation"
	RoutePairing = "/pairing"
	RouteBlock = "/block"
	RouteBan = "/ban"
	RouteAttribute = "/attribute"
	RouteParticipant = "/participant"
	RouteMessage = "/message"
	RouteReport = "/report"
	RouteFeedback = "/feedback"
	RouteSpecifyUser = "/:user_id"
	RouteSpecifyAttachment = "/:attachment_id"
	RouteSpecifyUniversity = "/:university_id"
	RouteSpecifyPairing = "/:pairing_id"
	RouteSpecifyConversation = "/:conversation_id"
	RouteSpecifyConversationParticipant = "/:conversation_participant_id"
	RouteSpecifyConversationMessage = "/:conversation_message_id"
	RouteSpecifyReport = "/:report_id"
	RouteSpecifyPairingAttribute = "/:pairing_attribute_id"
	RouteActionList = "_list"
)

var (
	AttachmentMediaTypeImage = "image"
	AttachmentMediaTypeVideo = "video"
	AttachmentUsedAsTypeMessage = "message"
	AttachmentUsedAsTypeAvatar = "avatar"
	AttachmentCategoryTypeUniversity = "university"
	AttachmentCategoryTypeUser = "user"
	AttachmentCategoryTypeConversation = "conversation"
)

var (
	ReportDecisionBanOffender = "ban_offender"
	ReportDecisionBanSender = "ban_sender"
	ReportDecisionNoAction = "no_action"
)