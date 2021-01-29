package requests

type BlockPut struct {
	UserId int64 `json:"user_id" json:"user_id" binding:"required,numeric"`
}

type BlockGet struct {
	UserId int64 `uri:"user_id" json:"user_id" binding:"required,numeric"`
}

type BlockList struct {
	ListRequest
}

type BlockDelete struct {
	UserId int64 `uri:"user_id" json:"user_id" binding:"required,numeric"`
}
