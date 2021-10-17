package app

import (
	"encoding/json"
	"net/http"

	"github.com/uhey22e/sqlboiler-restapi-example/restapi"
)

type ServerImpl struct{}

func (x *ServerImpl) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := ListUsers(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (x *ServerImpl) ListEvents(w http.ResponseWriter, r *http.Request) {
}

func Handler() http.Handler {
	x := &ServerImpl{}
	h := restapi.Handler(x)
	return h
}
