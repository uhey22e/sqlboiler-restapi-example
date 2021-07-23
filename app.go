package app

import (
	"encoding/json"
	"net/http"

	"github.com/uhey22e/sqlboiler-restapi-example/restapi"
)

type ServerImpl struct{}

func (x *ServerImpl) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := ListUsersAPI(ctx)
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

func Handler() http.Handler {
	x := &ServerImpl{}
	return restapi.Handler(x)
}
