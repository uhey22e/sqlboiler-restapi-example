package app

import (
	"bufio"
	"database/sql"
	"io"
	"os"

	"github.com/caarlos0/env/v6"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type dbConfig struct {
	URL string `env:"DB_URL" envDefault:"postgres://postgres:password@postgres/postgres?sslmode=disable"`
}

var (
	db  *sql.DB
	cfg *dbConfig
)

func ConfigureDatabase() {
	var err error
	db, err = sql.Open("pgx", cfg.URL)
	if err != nil {
		panic(err)
	}
}

func LoadSqlFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	queries := []string{}
	r := bufio.NewReader(f)
	for {
		q, err := r.ReadString(';')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		queries = append(queries, q)
	}
	return queries, nil
}

func init() {
	cfg = &dbConfig{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	ConfigureDatabase()
}
