package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/anne-markis/fermtrack/cli/client"
	"github.com/joho/godotenv"
)

type args struct {
	envFile string
}

func main() {
	ctx := context.Background()

	args := loadArgs()
	loadEnvVars(args.envFile)

	fermTracker := client.NewFermentationClient("http://0.0.0.0:8080") // TODO

	StartCLI(ctx, fermTracker)
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
		if strings.HasPrefix(arg, "env=") {
			vars := strings.Split(arg, "=")
			a.envFile = vars[1]
		}
	}
	return a
}
