package requests

type LogoutPost struct {
	AccountId int64 `json:"-" binding:"-"`
}

