package requests

type ListRequest struct {
	Limit int64 `form:"limit" json:"offset" binding:"omitempty,numeric"`
	AfterId int64 `form:"after_id" json:"after_id" binding:"omitempty,numeric"`
}
