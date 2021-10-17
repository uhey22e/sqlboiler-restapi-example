package app

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

type reqIdCtxKey struct{}

const reqIdHeaderKey = "X-Request-Id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), (reqIdCtxKey)(struct{}{}), id)
		hlog.FromRequest(r).UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("requestId", id)
		})
		w.Header().Add(reqIdHeaderKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
