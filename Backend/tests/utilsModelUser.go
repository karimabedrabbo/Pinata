package tests

import (
	"github.com/brianvoe/gofakeit/v4"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func GenerateUser() *dbmodel.User {
	return &dbmodel.User{
		BaseModel:   GenerateBaseModel(),
		AvatarId:    Increment(),
		Preferences: "",
		Name:        gofakeit.Name(),
		Biography:   gofakeit.Emoji(),
	}
}
