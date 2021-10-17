package main

import (
	"net/http"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	app "github.com/uhey22e/sqlboiler-restapi-example"
)

type Config struct {
	Debug bool `env:"DEBUG" envDefault:"false"`
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

	r := chi.NewRouter()
	h := app.Handler()
	r.Mount("/", h)
	http.ListenAndServe(":8080", r)
}
