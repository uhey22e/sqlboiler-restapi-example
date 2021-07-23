package app

import (
	"path/filepath"
	"testing"
)

func testConfigureDatabase(t testing.TB) {
	t.Helper()
	configureDatabase()

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
