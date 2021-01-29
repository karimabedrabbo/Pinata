package dbmodel

//marks ALL of a specific college's pairing options
type PairingMatches struct {
	BaseModel
	FromPairingParticipantId int64 `gorm:"index;not null" json:"from_pairing_participant_id"`
	ToPairingParticipantId string `gorm:"index;not null" json:"to_pairing_participant_id"`
}

