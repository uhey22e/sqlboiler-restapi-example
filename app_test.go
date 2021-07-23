package app

import (
	"bytes"
	"encoding/json"
	"testing"
)

func logJson(t testing.TB, v interface{}) {
	t.Helper()
	w := bytes.NewBuffer(nil)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		t.Error(err)
		return
	}
	t.Logf("\n%s", w.String())
}
