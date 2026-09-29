package midelware

import (
	"net/http"

	"github.com/google/uuid"
)

const requestIdKey = "X-Request-ID"

func RequestId() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIdKey)

			if requestId == "" {
				requestId = uuid.NewString()
			}

			r.Header.Set(requestIdKey, requestId)
			w.Header().Set(requestIdKey, requestId)

			next.ServeHTTP(w, r)
		})
	}
}
