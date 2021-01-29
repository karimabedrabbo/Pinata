package requests

type ChangeEmailPost struct {
	UniversityId int64 `json:"-" binding:"-"`
	AccountId int64 `json:"-" binding:"-"`
	NewEmail string `json:"new_email" binding:"required,email,endswith=.edu"`
}
