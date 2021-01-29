package managers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	k "github.com/karimabedrabbo/eyo/api/constants"
	"github.com/karimabedrabbo/eyo/api/database"
	"github.com/karimabedrabbo/eyo/api/handlers"
)


func SetupRequestHandler(c *gin.Context) *handlers.RhEnv {
	// begin a transaction so that the entire request is all or nothing in the database
	tx := GetDatabase().GormClient.BeginTx(c, &sql.TxOptions{})
	return &handlers.RhEnv{
		E:        &database.DbEnv{Tx:&tx},
		C:        c,
		Mail:     GetMail(),
		Sanitize: GetSanitize(),
		Storage:  GetStorage(),
		Auth:     GetAuthentication(),
	}
}

func InitRequestHandler(c *gin.Context) {
	c.Set(k.RequestHandlerKey, SetupRequestHandler(c))
}