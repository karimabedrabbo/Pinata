package dbmodel

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type BaseModel struct {
	Id        int64       `gorm:"primary_key" json:"id"`
	CreateAt int64      `json:"create_at"` //not "created" because gorm messes things up with int64
	UpdateAt int64      `json:"update_at"` //not "updated" because gorm messes things up with int64
	DidPrepare bool		`gorm:"-" json:"-"`  //this is to check that before we save anything we have called
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}


type DatabaseModel interface {
	Prepare()
}

func (m *BaseModel) Prepare() {
	m.DidPrepare = true
}

func (m *BaseModel) BeforeSave (scope *gorm.Scope) error {
	var err error

	if m.DidPrepare == false {
		return apperror.PrepareMarkerMissing
	}

	if m.UpdateAt == 0 {
		if err = m.setUpdateAt(scope); err != nil {
			return err
		}
	}

	if err = m.setCreateAt(scope); err != nil {
		return err
	}
	return nil
}

func (m *BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	var err error
	if m.DidPrepare == false {
		return apperror.PrepareMarkerMissing
	}
	if err = m.setUpdateAt(scope); err != nil {
		return err
	}
	return nil
}

func (m *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	var err error

	if m.DidPrepare == false {
		return apperror.PrepareMarkerMissing
	}

	if m.UpdateAt == 0 {
		if err = m.setUpdateAt(scope); err != nil {
			return err
		}
	}

	if err = m.setCreateAt(scope); err != nil {
		return err
	}
	return nil
}

func (m *BaseModel) setUpdateAt(scope *gorm.Scope) error {
	return scope.SetColumn("update_at", time.Now().Unix())
}

func (m *BaseModel) setCreateAt(scope *gorm.Scope) error {
	return scope.SetColumn("create_at", time.Now().Unix())
}
