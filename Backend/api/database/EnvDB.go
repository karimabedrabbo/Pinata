package database

import "github.com/jinzhu/gorm"

type DbEnv struct {
	//double pointer here due to gorm method chaining (preserve outer pointer so it's not being modified)
	Tx **gorm.DB
}


func (e *DbEnv) GetTx() *gorm.DB {
	return *e.Tx
}

func (e *DbEnv) SetTx(tx *gorm.DB) {
	*e.Tx = tx
}

func (e *DbEnv) IsRecordNotFound(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

