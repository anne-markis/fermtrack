package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/anne-markis/fermtrack/answer"

	"github.com/joho/godotenv"
)

type args struct {
	cheapmode bool
	envFile   string
}

func main() {
	ctx := context.Background()

	args := loadArgs()
	loadEnvVars(args.envFile)

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

	StartCLI(ctx, aiClient)
}

func loadEnvVars(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func loadArgs() args {
	a := args{
		envFile: "../.env",
	}
	for _, arg := range os.Args {
		if arg == "cheap" {
			a.cheapmode = true
			continue
		}
		if strings.HasPrefix(arg, "env=") {
			vars := strings.Split(arg, "=")
			a.envFile = vars[1]
		}
	}
	return a
}
