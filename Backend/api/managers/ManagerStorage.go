package managers

import (
	gstorage "cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/karimabedrabbo/eyo/api/apperror"
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/models/response"
	"net/http"
	"time"
)

type Storage struct {
	StorageClient *gstorage.Client
}

var storage *Storage

func SetupStorage() *Storage {
	//we dont need the actual client for now
	return &Storage{}
}

func InitStorage() {
	storage = SetupStorage()
}

func GetStorage() *Storage {
	return storage
}

func (s *Storage) GetSignedUrl(method string, mediaType string, encodingType string, md5 string, attachmentId int64, tokenUuid uuid.UUID) (*response.SignedUrlResponse, error) {
	var err error
	var url string

	if method != http.MethodPut && method != http.MethodGet && method != http.MethodDelete {
		return nil, apperror.StorageInvalidMethod
	}

	contentType := mediaType + "/" + encodingType

	expiresAt := time.Now().Add(time.Minute)
	url, err = gstorage.SignedURL(k.StorageBucketName, apputils.IdToString(attachmentId),
		&gstorage.SignedURLOptions{
			GoogleAccessID:apputils.GetGoogleCloudServiceEmail(),
			PrivateKey:apputils.GetGoogleCloudServicePrivateKey(),
			Method: method,
			MD5: md5,
			ContentType: contentType,
			Expires: expiresAt,
			Headers: []string{"x-goog-token-uuid", "x-goog-attachment-id"},
		})
	if err != nil {
		return nil, err
	}

	signedUrlResponse := &response.SignedUrlResponse{
		AttachmentId: attachmentId,
		TokenUuid: tokenUuid,
		SignedUrl: url,
		ExpiresAt: expiresAt.Unix(),
		Method: method,
	}

	return signedUrlResponse, nil
}