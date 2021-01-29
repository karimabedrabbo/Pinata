package dbmodel

type PairingPreference struct {
	BaseModel
	PairingParticipantsId int64 `gorm:"index;not null" json:"pairing_participant_id"`
	PairingAttributeId int64 `gorm:"not null" json:"pairing_attribute_id"`
}
