package main

import (
	"eniqlo/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	s := server.NewServer()

	log.Fatal(s.Run())

}
