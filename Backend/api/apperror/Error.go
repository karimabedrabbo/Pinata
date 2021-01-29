package apperror

import "errors"

var (
	AuthenticationUserUnverified = errors.New("unverified user attempted to authenticate")
	AuthenticationUserBanned = errors.New("banned user attempted to authenticate")
	PrepareMarkerMissing = errors.New("prepare was not called or not set on model")
	PrepareIdentityUninitialized = errors.New("request did not have identity to provide to model")
	PrepareMissingAttribute = errors.New("request missing a required attribute")
	DatabaseParsingError = errors.New("encountered an error parsing input")
	DatabaseIdentityUninitialized = errors.New("using model without an identity where one is required")
	DatabaseModelUninitialized = errors.New("using an uninitialized model")
	RequestWindowExpired = errors.New("missed the available window to make this request")
	RequestAlreadyUsed = errors.New("already issued this one time use request")
	RequestInvalidCredentials = errors.New("invalid credentials supplied in request")
	RequestTokenIdentityInvalid = errors.New("provided invalid identity in the token to the handler")
	RequestTokenIdentityMissing = errors.New("did not provide an identity in the token to the handler")
	RequestInvalidForm = errors.New("request interface of an invalid form")
	RequestNotSupported = errors.New("request not supported")
	RequestForbidden = errors.New("request forbidden")
	StorageInvalidMethod =  errors.New("attempted to access an invalid method")
)