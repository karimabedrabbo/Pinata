package tests

import (
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
)

func GenerateBan(accountId int64) *dbmodel.Ban {
	return &dbmodel.Ban{
		BaseModel:       GenerateBaseModel(),
		ToAccountId: accountId,
		IssuerAccountId:  Increment(),
	}
}


