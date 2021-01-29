package tests

import (
	"github.com/google/uuid"
	"github.com/karimabedrabbo/eyo/api/apputils"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"log"
	"time"
)

func GenerateVerify(accountId int64, verified bool) (*dbmodel.Verify, string) {
	var err error

	tokenUuid := uuid.New().String()
	var tokenUuidHash string
	if tokenUuidHash, err = apputils.DerivePasswordHash(tokenUuid); err != nil {
		log.Fatalf("could not hash generated passowrd: %v", err)
	}

	var usedTime int64 = 0
	if verified {
		usedTime = time.Now().Unix()
	}

	return &dbmodel.Verify{
		BaseModel:       GenerateBaseModel(),
		AccountId:       accountId,
		HashedTokenUuid: tokenUuidHash,
		UsedAt:          usedTime,
		ExpiresAt:       time.Now().Add(time.Hour * 24).Unix(),
	}, tokenUuid
}

