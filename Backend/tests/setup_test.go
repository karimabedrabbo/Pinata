package tests

import (
	fuzz "github.com/google/gofuzz"
	"github.com/joho/godotenv"
	"github.com/karimabedrabbo/eyo/api/managers"
	"github.com/karimabedrabbo/eyo/api/server"
	"log"
	"os"
	"testing"
)

var s *TestSuite
var CREDENTIALS_REPEAT_ITER int = 10
var REPEAT_ITER int = 100
var GEN_COUNT int = 100

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error getting env, %v", err)
	}

	server.Initialize()
	s = &TestSuite{
		router:   managers.GetRouter(),
		db:       managers.GetDatabase(),
		auth:     managers.GetAuthentication(),
		sanitize: managers.GetSanitize(),
		fuzzer:   fuzz.New(),
	}
	CleanDatabase(s.db)
	os.Exit(m.Run())
}

