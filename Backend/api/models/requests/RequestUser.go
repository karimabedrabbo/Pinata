package requests

import (
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type UserProfileGet struct {
	UserId int64 `json:"-" binding:"-"`
}

type UserAvatarGet struct {
	UserId int64 `json:"-" binding:"-"`
}

type UserProfilePost struct {
	UserId int64 `json:"-" binding:"-"`
	Name string `json:"name" binding:"omitempty,max=255"`
	Biography 		string 	`json:"biography" binding:"omitempty,max=600"`
	Preferences string `json:"preferences,string" binding:"omitempty,json"`
}

type UserAvatarPost struct {
	UserId int64 `json:"-" binding:"-"`
	*ImageAttachmentCreatable
}

type DifferentUserProfileGet struct {
	UserId int64 `uri:"user_id" json:"user_id" binding:"required,numeric"`
}

type DifferentUserAvatarGet struct {
	UserId int64 `uri:"user_id" json:"user_id" binding:"required,numeric"`
}



