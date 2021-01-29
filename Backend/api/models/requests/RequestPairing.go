package requests

type PairingId struct {
	PairingId int64 `uri:"pairing_id" json:"pairing_id" binding:"required,numeric"`
}

type PairingUpdatable struct {
	Preferences []string `json:"preferences" binding:"omitempty,dive,unique"`
}

type PairingCreateable struct {
	PairingUpdatable
}

type PairingGet struct {
	PairingId
}

type PairingPut struct {
	PairingCreateable
}

type PairingPost struct {
	PairingId
	PairingUpdatable
}

type PairingDelete struct {
	PairingId
}