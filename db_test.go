package app

import (
	"database/sql"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
)

func testConfigureDatabase(t testing.TB) {
	t.Helper()
	var err error
	uri := "postgres://postgres:password@postgres/postgres?sslmode=disable"
	c, _ := pgx.ParseConfig(uri)
	c.Logger = zerologadapter.NewLogger(log.Logger)
	c.LogLevel = pgx.LogLevelInfo

	db, err = sql.Open("pgx", stdlib.RegisterConnConfig(c))
	if err != nil {
		t.Fatal(err)
	}

	queries, err := LoadSqlFile(filepath.Join("testdata", "seeds.sql"))
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	for _, q := range queries {
		_, err = tx.Exec(q)
		if err != nil {
			tx.Rollback()
			t.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadSqlFile(t *testing.T) {
	got, err := LoadSqlFile(filepath.Join("testdata", "seeds.sql"))
	if err != nil {
		t.Errorf("LoadSqlFile() error = %v", err)
		return
	}
	for _, q := range got {
		t.Log(q)
	}
}
