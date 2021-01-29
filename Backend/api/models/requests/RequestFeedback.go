package requests

type FeedbackPut struct {
	UserId int64 `json:"-" binding:"-"`
	Title string `json:"title" binding:"required"`
	Comment string `json:"comment" binding:"omitempty"`
}
