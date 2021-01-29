package requests

type ChangePasswordPost struct {
	AccountId int64 `json:"-" binding:"-"`
	OldPassword string `json:"old_password" binding:"required,min=8"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}
