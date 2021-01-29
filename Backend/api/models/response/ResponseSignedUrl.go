package response

import "github.com/google/uuid"

type SignedUrlResponse struct {
	AttachmentId int64 `json:"attachment_id"`
	TokenUuid uuid.UUID `json:"token_uuid"`
	SignedUrl string `json:"signed_url"`
	ExpiresAt int64 `json:"expires_at"`
	Method string `json:"method"`
}
