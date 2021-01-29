package response

type PayloadTokenResponse struct {
	Token string `json:"token"`
	ExpiresAt int64 `json:"expires_at"`
}
