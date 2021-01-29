package requests

type SignupPut struct {
	Name string `json:"name" binding:"required,max=255"`
	Email string `json:"email" binding:"required,email,endswith=.edu"`
	Password string `json:"password" binding:"required,min=8"`
}