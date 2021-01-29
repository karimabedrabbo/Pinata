package dbmodel

type PairingBoard struct {
	BaseModel
	UniversityId int64 `gorm:"unique_index" json:"university_id"`
	University   *University `gorm:"foreignkey:UniversityId"`
	Name         string `gorm:"size:255" json:"name"`
	Attributes     *[]PairingAttribute `gorm:"foreignkey:PairingAttributeId"`
	Participants     *[]PairingParticipant `gorm:"foreignkey:PairingParticipantId"`
}

