package tests

import (
	"github.com/karimabedrabbo/eyo/api/managers"
	"github.com/karimabedrabbo/eyo/api/models/response"
	"log"
)

func TokenGenerator(auth *managers.Auth, userId int64, accountId int64, universityId int64, role string, emailVerified bool, universityVerified bool) string {
	payload := &response.PayloadClaims{
		UserId:             userId,
		AccountId:          accountId,
		UniversityId:       universityId,
		Role:               role,
		EmailVerified:      emailVerified,
		UniversityVerified: universityVerified,
	}

	token, _, err := auth.GinJwtClient.TokenGenerator(payload)
	if err != nil {
		log.Fatalf("could not generate token: %v", err)
	}

	return token
}

func TokenGeneratorAnonymous(auth *managers.Auth, userId int64, universityId int64) string {
	return TokenGenerator(auth, userId, 0, universityId, "anonymous", false, false)
}

func TokenGeneratorUnverifiedUser(auth *managers.Auth, userId int64, accountId int64, universityId int64) string {
	return TokenGenerator(auth, userId, accountId, universityId, "user", false, false)
}

func TokenGeneratorUser(auth *managers.Auth, userId int64, accountId int64, universityId int64) string {
	return TokenGenerator(auth, userId, accountId, universityId, "user", true, true)
}

func TokenGeneratorSysadmin(auth *managers.Auth) string {
	return TokenGenerator(auth, 1, 1, 1, "sysadmin", true, true)
}

