package requests

type LoginAnonymousPost struct {
	UniversityId int64 `json:"university_id" binding:"required,numeric"`
}

type LoginPost struct {
	Email string `json:"email" binding:"required,email,endswith=.edu"`
	Password string `json:"password" binding:"required,min=8"`
}