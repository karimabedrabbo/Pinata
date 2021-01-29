package main

import (
	"github.com/joho/godotenv"
	"github.com/karimabedrabbo/eyo/api/apputils"
	"github.com/karimabedrabbo/eyo/api/server"
	"log"
)

func main() {


	loadEnv()

	server.Initialize()

	server.Serve()

}

func loadEnv() {
	if apputils.GetAppContextIsLocal() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("error getting env %v", err)
		}
	}
}