package apputils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordIsHashed(hash string) error {
	_, err := bcrypt.Cost([]byte(hash))
	return err
}

func DerivePasswordHash(password string) (string, error) {
	var cost int = 14
	if !GetAppEnvIsProduction() {
		cost = 1
	}
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GetUUID() (uuid.UUID, error) {
	var newUUID uuid.UUID
	var err error
	newUUID, err = uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}
	return newUUID, nil
}

