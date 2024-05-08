package main

import (
	"eniqlo/server"
	"log"

	"github.com/joho/godotenv"
)

// @title Swagger Example API
// @version 1.0

// @BasePath /v1
func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	s := server.NewServer()

	log.Fatal(s.Run())

}
