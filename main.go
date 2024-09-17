package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvVars()

	// TODO
	// cli.StartCLI()
}

func loadEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
