package dbmodel

import (
	"github.com/jinzhu/gorm"
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/models/requests"
)

type Account struct {
	BaseModel
	UserId   int64  `gorm:"index;not null" json:"user_id"`
	Email   string `gorm:"size:255;unique_index;not null" json:"email"`
	PasswordHash  string `gorm:"size:100;not null" json:"-"`
	Role 		string `gorm:"size:100;default:'account'" json:"role"`
}


func (account *Account) PrepareChangeEmailPost(r *requests.ChangeEmailPost) error {
	if r.AccountId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	account.Id = r.AccountId
	account.Email = r.NewEmail
	account.Prepare()
	return nil
}

func (account *Account) PrepareChangePasswordPost(r *requests.ChangePasswordPost) error {
	if r.AccountId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	account.Id = r.AccountId
	passHash, err := apputils.DerivePasswordHash(r.NewPassword)
	if err != nil {
		return err
	}
	account.PasswordHash = passHash
	account.Prepare()
	return nil
}

func (account *Account) PreparePasswordResetPost(r *requests.PasswordResetPost) error {
	if r.AccountId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	account.Id = r.AccountId
	passHash, err := apputils.DerivePasswordHash(r.NewPassword)
	if err != nil {
		return err
	}
	account.PasswordHash = passHash
	account.Prepare()
	return nil
}

func (account *Account) PrepareSignupPut(r *requests.SignupPut) error {
	passHash, err := apputils.DerivePasswordHash(r.Password)
	if err != nil {
		return err
	}
	account.PasswordHash = passHash
	account.Email = r.Email
	account.Prepare()
	return nil
}


func (account *Account) BeforeSave (scope *gorm.Scope) error {
	if err := apputils.CheckPasswordIsHashed(account.PasswordHash); err != nil  {
		return err
	}
	return nil
}

func (account *Account) BeforeCreate (tx *gorm.DB) (err error) {
	if err := apputils.CheckPasswordIsHashed(account.PasswordHash); err != nil  {
		return err
	}
	return
}

func (account *Account) BeforeUpdate (tx *gorm.DB) (err error) {
	if err := apputils.CheckPasswordIsHashed(account.PasswordHash); err != nil {
		return err
	}
	return
}