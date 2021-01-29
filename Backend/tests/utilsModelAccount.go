package tests

import (
	"github.com/brianvoe/gofakeit/v4"
	"github.com/karimabedrabbo/eyo/api/apputils"
	k "github.com/karimabedrabbo/eyo/api/constants"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"log"
)

func GenerateAccounts(count int) []*dbmodel.Account {
	accounts := make([]*dbmodel.Account, count)
	for i := 0; i < count; i++ {
		accounts[i], _ = GenerateAccount("", "")
	}
	return accounts
}

func GenerateEmail(universityDomain string) string {
	return IncrementString() + "@" + universityDomain
}

func GenerateAccount(email string, role string) (*dbmodel.Account, string) {
	var err error

	if role == "" {
		role = k.AccountRoleUser
	}

	password := gofakeit.Password(true, true, true, true, true, 8)
	var passHash string
	if passHash, err = apputils.DerivePasswordHash(password); err != nil {
		log.Fatalf("could not hash generated passowrd: %v", err)
	}
	return &dbmodel.Account{
		BaseModel:    GenerateBaseModel(),
		UserId:       Increment(),
		Email:        email,
		PasswordHash: passHash,
		Role:         role,
	}, password
}
