package dbmodel

import (
	"github.com/karimabedrabbo/eyo/api/apperror"
	requests "github.com/karimabedrabbo/eyo/api/models/requests"
)

type University struct {
	BaseModel
	Name         string        `gorm:"size:255;index;not null" json:"name"`
	Alias        string        `gorm:"size:255;index" json:"alias"`
	Domain  string        `gorm:"size:255;index;not null;unique" json:"domain"`
	Users        *[]User       `gorm:"foreignkey:UserId"`
	Lat          float32       `gorm:"type:decimal(10,8)" json:"lat"`
	Lng          float32       `gorm:"type:decimal(11,8)" json:"lng"`
	AvatarId     int64         `gorm:"index" json:"avatar_id"`
	Size         int          `gorm:"index" json:"size"`
	PairingBoard *PairingBoard `gorm:"foreignkey:PairingBoardId"`
}

func (univ *University) PrepareUniversityPut(r *requests.UniversityPut) error {
	univ.Name = r.Name
	univ.Alias = r.Alias
	univ.Domain = r.Domain
	univ.Lat = r.Lat
	univ.Lng = r.Lng
	univ.Size = r.Size
	univ.Prepare()
	return nil
}

func (univ *University) PrepareUniversityPost(r *requests.UniversityPost) error {
	if r.UniversityId == 0 {
		return apperror.PrepareIdentityUninitialized
	}
	univ.Id = r.UniversityId
	univ.Name = r.Name
	univ.Alias = r.Alias
	univ.Size = r.Size
	univ.Prepare()
	return nil
}