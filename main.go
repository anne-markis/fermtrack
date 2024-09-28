package main

import (
	"cmp"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/anne-markis/fermtrack/server"

	"github.com/joho/godotenv"
)

type args struct {
	cheapmode bool
	envFile   string
}

func main() {
	// ctx := context.Background()

	args := loadArgs()
	loadEnvVars(args.envFile)

	// var aiClient answer.AnsweringClient

	// if args.cheapmode {
	// 	aiClient = answer.CheapClient{}
	// } else {
	// 	var err error
	// 	aiClient, err = answer.InitClient()
	// 	if err != nil {
	// 		log.Fatalf("failed to  load open ai client: %s", err)
	// 	}
	// }

	ftServer := server.FermtrackServer{}

	// start server
	mux := http.NewServeMux()
	mux.HandleFunc("GET /list/", ftServer.ListProjectsHandler)
	mux.HandleFunc("POST /edit/", ftServer.EditProjectHandler)

	// start up server
	go func() {
		port := cmp.Or(os.Getenv("SERVERPORT"), "9020")
		address := "localhost:" + port
		log.Printf("listening on %s\n", port)
		if err := http.ListenAndServe(address, mux); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	// // shut down gracefully
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	<-ctx.Done()
	// 	shutdownCtx := context.Background()
	// 	shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
	// 	defer cancel()
	// 	if err := http.Shutdown(shutdownCtx); err != nil {
	// 		fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
	// 	}
	// }()
	// wg.Wait()

	// cli.StartCLI(ctx, aiClient) // send server info?
}

func loadEnvVars(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func loadArgs() args {
	a := args{
		envFile: ".env",
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
