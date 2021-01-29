package dbmodel

type PairingParticipant struct {
	BaseModel
	PairingBoardId int64 `gorm:"index;not null" json:"pairing_board_id"`
	UserId int64 `gorm:"unique_index;not null" json:"user_id"`
	Lat                 float32 `gorm:"type:decimal(10,8)" json:"lat"`
	Lng                 float32 `gorm:"type:decimal(11,8)" json:"lng"`
	Preferences *[]PairingPreference `gorm:"foreignkey:PairingParticipantPreferenceId"`
}
