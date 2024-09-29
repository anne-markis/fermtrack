package main

import (
	"context"
	"flag"

	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/anne-markis/fermtrack/internal/app"
	"github.com/anne-markis/fermtrack/internal/handlers"
	"github.com/anne-markis/fermtrack/internal/repository"
	"github.com/anne-markis/fermtrack/internal/router"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

type args struct {
	cheapmode bool
	envFile   string
}

func main() {

	args := loadArgs()
	loadEnvVars(args.envFile)

	// TODO put connection something else
	db, err := sql.Open("mysql", "root:s3CrEt@tcp(mysql:3306)/fermtrack?parseTime=true")
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	defer db.Close()

	for {
		if err := db.Ping(); err != nil {
			log.Info().Err(err).Msg("DB not ready, retrying")
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	log.Info().Msg("db ready, connected")

	if err := goose.SetDialect("mysql"); err != nil {
		log.Error().Err(err).Msg("Failed to set dialect")
	}
	if err := goose.Up(db, os.Getenv("GOOSE_MIGRATION_DIR")); err != nil {
		log.Error().Err(err).Msg("Error running migrations")
		return
	}

	repo := repository.NewMySQLFermentationRepository(db)
	fermService := app.NewFermentationService(repo)
	fermHandler := handlers.NewFermentationHandler(fermService)

	r := router.NewRouter(fermHandler)

	// middleware
	r.Use(loggingMiddleware)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080", // TODO
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error().Msg(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Info().Msg("shutting down")
	os.Exit(0)
}

func loadEnvVars(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal().Msg(err.Error())
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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
