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
	_ = godotenv.Load(".env")

	s := server.NewServer()

	log.Fatal(s.Run())

}
