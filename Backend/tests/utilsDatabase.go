package tests

import (
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/managers"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"log"
)

var inc int64 = 1

func Increment() int64 {
	temp := inc
	inc += 1
	return temp
}

func IncrementString() string {
	temp := inc
	inc += 1
	tempString := apputils.IdToString(temp)
	return tempString
}

func Seed(db *managers.Database, model interface{}) {
	var err error
	if err = db.GormClient.Create(model).Error; err != nil {
		log.Fatalf("could not seed database: %v", err)
	}
}

func CleanDatabase(db *managers.Database) {
	var err error
	for _, table := range db.GetAllTables() {
		if err = db.GormClient.DropTableIfExists(table).Error; err != nil {
			log.Fatalf("could not drop table: %v", err)
		}
	}
}

func GenerateBaseModel() dbmodel.BaseModel {
	//only set the did prepare as it's required on update and create
	//the rest will be filled in when we seed it into the database
	return dbmodel.BaseModel{
		DidPrepare: true,
	}
}

func SetupDatabase(db *managers.Database) {
	//renew the tables fresh
	db.MigrateTables()
}

func RefreshDatabase(db *managers.Database) {
	CleanDatabase(db)
	SetupDatabase(db)
}