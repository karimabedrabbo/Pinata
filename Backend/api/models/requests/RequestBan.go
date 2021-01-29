package requests

type BanPut struct {
	ToAccountId int64 `json:"to_account_id" binding:"required,numeric"`
	IssuerAccountId int64 `json:"-" binding:"-"`
}

type BanGet struct {
	ToAccountId int64 `json:"to_account_id" binding:"required,numeric"`
}

type BanDelete struct {
	ToAccountId int64 `json:"to_account_id" binding:"required,numeric"`
}
