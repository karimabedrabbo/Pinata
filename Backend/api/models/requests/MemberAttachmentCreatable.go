package requests

import "github.com/google/uuid"

type ImageAttachmentCreatable struct {
	TokenUuid uuid.UUID `json:"-" binding:"-"`
	Md5  string  `json:"md5" binding:"required,hexadecimal"`
	EncodingType string `json:"encoding_type" binding:"required,oneof=png jpeg"`
}

type VideoAttachmentCreatable struct {
	TokenUuid uuid.UUID `json:"-" binding:"-"`
	Md5  string  `json:"md5" binding:"required,hexadecimal"`
	EncodingType string `json:"encoding_type" binding:"required,eq=H264"`
}