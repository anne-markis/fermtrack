package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anne-markis/fermtrack/internal/app"
	"github.com/anne-markis/fermtrack/internal/app/ai"
	"github.com/anne-markis/fermtrack/internal/app/domain"
	"github.com/anne-markis/fermtrack/internal/config"
	"github.com/anne-markis/fermtrack/internal/handlers"

	"github.com/anne-markis/fermtrack/internal/router"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

func main() {

	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)
	db, err := sql.Open("mysql", dsn)
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

	aiClient, err := ai.InitClient()
	if err != nil {
		log.Error().Err(err).Msg("failed to setup AI client")
	}

	fermRepo := domain.NewMySQLFermentationRepository(db)
	userRepo := domain.NewMySQLUserRepository(db)

	authService := app.NewAuthService(userRepo)
	fermService := app.NewFermentationService(fermRepo, aiClient)

	fermHandler := handlers.NewFermentationHandler(fermService)
	authHandler := handlers.NewAuthHandler(authService, userRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	r := router.NewRouter(fermHandler, authHandler, userHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", cfg.Server.Port),
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
