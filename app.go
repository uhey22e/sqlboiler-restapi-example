package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/uhey22e/sqlboiler-restapi-example/restapi"
)

type ServerImpl struct{}

func (x *ServerImpl) ListUsers(w http.ResponseWriter, r *http.Request) {
	items, err := ListUsers(r.Context())
	if err != nil {
		hlog.FromRequest(r).Error().Err(err).Msg("Error occurred.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeJsonResponse(w, items)
}

func (x *ServerImpl) ListEvents(w http.ResponseWriter, r *http.Request) {
	items, err := ListEvents(r.Context())
	if err != nil {
		hlog.FromRequest(r).Error().Err(err).Msg("Error occurred.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeJsonResponse(w, items)
}

func (x *ServerImpl) ListPopularEvents(w http.ResponseWriter, r *http.Request) {
	items, err := ListPopularEvents(r.Context())
	if err != nil {
		hlog.FromRequest(r).Error().Err(err).Msg("Error occurred.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeJsonResponse(w, items)
}

func writeJsonResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Handler() http.Handler {
	r := chi.NewMux()
	r.Use(hlog.NewHandler(log.Logger))
	r.Use(RequestID)
	r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("Response served.")
	}))
	r.Use(hlog.RemoteAddrHandler("ip"))

	r.Get("/system/health", func(w http.ResponseWriter, r *http.Request) {
		writeJsonResponse(w, struct{}{})
	})

	x := &ServerImpl{}
	h := restapi.HandlerFromMux(x, r)
	return h
}
