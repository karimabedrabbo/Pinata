package managers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/karimabedrabbo/eyo/api/apputils"
	dbmodel "github.com/karimabedrabbo/eyo/api/models/database"
	"log"
)

type Database struct {
	//double pointer here due to gorm method chaining (preserve outer pointer so it's not being modified)
	GormClient *gorm.DB
}

var db *Database

func SetupDatabase() *Database {
	var err error


	dbHost := apputils.GetDatabaseHost()
	dbPort := apputils.GetDatabasePort()
	dbUser := apputils.GetDatabaseUser()
	dbName := apputils.GetDatabaseName()
	dbPassword := apputils.GetDatabasePassword()

	dbUrlNoPassword := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName)
	dbUrlFull := fmt.Sprintf("%s password=%s", dbUrlNoPassword, dbPassword)

	fmt.Printf("attempting to connect to postgres:\nconnection config: %s\n", dbUrlNoPassword)

	gormClient, err := gorm.Open("postgres", dbUrlFull)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	fmt.Printf("successfully connected to postgres\n")

	//do not use plural table names
	gormClient.SingularTable(true)

	gormClient.BlockGlobalUpdate(true)


	tempDb := &Database{
		GormClient: gormClient,
	}

	tempDb.MigrateTables()

	return tempDb
}

func InitDatabase() {
	db = SetupDatabase()
}

func GetDatabase() *Database {
	return db
}


func (db *Database) GetAllTables() []interface{} {
	tables := []interface{}{
		&dbmodel.BaseModel{},
		&dbmodel.User{},
		&dbmodel.University{},
		&dbmodel.Account{},
		&dbmodel.Verify{},
		&dbmodel.Ban{},
		&dbmodel.PasswordReset{},
		&dbmodel.Report{},
	}
	return tables
}

func (db *Database) MigrateTables() {
	for _, table := range db.GetAllTables() {
		db.GormClient.AutoMigrate(table)
	}
}
