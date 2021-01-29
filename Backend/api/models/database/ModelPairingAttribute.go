package dbmodel

//marks ALL of a specific college's pairing boards options
type PairingAttribute struct {
	BaseModel
	PairingBoardId int64 `gorm:"index;not null" json:"pairing_board_id"`
	Attribute string `gorm:"size:100;not null" json:"attribute"`
}
