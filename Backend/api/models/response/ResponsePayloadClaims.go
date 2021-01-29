package response


type PayloadClaims struct {
	UserId int64 `json:"user_id"`//here for every user
	AccountId int64 `json:"account_id"`//only here if user has an account (verified or unverified)
	Role string `json:"role"`//prefilters access to certain endpoints
	EmailVerified bool `json:"email_verified"` //required to know that user with an account actually verified they're from university they claim
	UniversityId int64 `json:"university_id"`//we need this here (and can't just refer to database) because this stores an anonymous user's university
	UniversityVerified bool `json:"university_verified"`//required to know that user with an account actually verified they're from university they claim
}
