package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/anne-markis/fermtrack/answer"
	"github.com/joho/godotenv"
)

type args struct {
	cheapmode bool
}

func main() {
	ctx := context.Background()

	loadEnvVars()
	args := loadArgs()

	var aiClient answer.AnsweringClient

	if args.cheapmode {
		aiClient = answer.CheapClient{}
	} else {
		var err error
		aiClient, err = answer.InitClient()
		if err != nil {
			log.Fatalf("failed to  load open ai client: %s", err)
		}
	}

	answer, err := aiClient.AskQuestion(ctx, "test wine") // TODO
	if err != nil {
		fmt.Println("failed to get answer")
	}
	fmt.Println(answer)

	// TODO need to rethink the bubbletea aspect
	// cli.StartCLI()
}

func loadEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func loadArgs() args {
	a := args{}
	for _, arg := range os.Args {
		if arg == "cheap" {
			a.cheapmode = true
		}
	}
	return a
}
