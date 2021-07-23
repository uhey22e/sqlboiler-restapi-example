package main

import (
	"net/http"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	app "github.com/uhey22e/sqlboiler-restapi-example"
)

type Config struct {
	Debug       bool   `env:"DEBUG" envDefault:"false"`
	DatabaseURL string `env:"DB_URL" envDefault:"postgres://postgres:password@postgres/postgres?sslmode=disable"`
}

var (
	cfg *Config
)

func configureLogger() {
	if cfg.Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func main() {
	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	configureLogger()

	r := chi.NewMux()
	r.Use(middleware.Logger)

	r.Use(hlog.NewHandler(log.With().Str("variant", "access").Logger()))
	r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("response served")
	}))

	app := app.Handler()
	r.Mount("/", app)
	http.ListenAndServe(":8080", app)
}
