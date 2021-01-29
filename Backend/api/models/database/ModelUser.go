package dbmodel

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/models/requests"
)

type User struct {
	BaseModel
	AvatarId   	int64  `gorm:"index" json:"avatar_id"`
	Preferences string `gorm:"text" json:"preferences"`
	Name        string `gorm:"size:255" json:"name"`
	Biography 		string 	`gorm:"text" json:"biography"`
}

func (user *User) PrepareUserPost(r *requests.UserProfilePost) error {
	if r.UserId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	user.Id = r.UserId
	user.Name = r.Name
	user.Biography = r.Biography
	user.Prepare()
	return nil
}

func (user *User) PrepareSignupPut(r *requests.SignupPut) error {
	user.Name = r.Name
	user.Prepare()
	return nil
}

func (user *User) PrepareLoginAnonymousPost(r *requests.LoginAnonymousPost) error {
	user.Prepare()
	return nil
}

//gorm hooks
func (user *User) TableName() string {
	return "users"
}