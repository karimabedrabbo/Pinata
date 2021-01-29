package server

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/managers"
	"log"
	"math/rand"
	"time"
)

func Initialize() {


	fmt.Printf("Server initializing:\napp env: %s\napp context: %s\n", apputils.GetAppEnv(), apputils.GetAppContext())

	//seed the server
	var random int64 = 83019823405
	rand.Seed(time.Now().UnixNano() + random)

	managers.InitDatabase()

	managers.InitRedis()

	managers.InitAuthentication()

	managers.InitIpFilter()

	managers.InitSanitize()

	managers.InitStorage()

	//managers.InitMessageQueue()

	managers.InitRateLimit()

	managers.InitMail()

	managers.InitRouter()

}

func Serve() {
	apiPort := fmt.Sprintf(":%s", apputils.GetApiPort())
	fmt.Printf("Listening to port %s", apiPort)
	log.Fatal(managers.GetRouter().Engine.Run(apiPort))
}




