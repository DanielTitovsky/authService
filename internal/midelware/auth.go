package midelware

import (
	usecases "authService/internal/useCases"
	jwt_utils "authService/internal/utils/jwt"
	transport_http_utils "authService/internal/utils/transport/http"
	"net/http"
)

func AuthMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responceHandler := transport_http_utils.NewHTTPResponceHandler(w)

			coockie, err := r.Cookie(usecases.AccessTokenKey)

			if err != nil {
				responceHandler.Error(http.StatusForbidden, "Undefined token string", "Login to continue")
				return
			}

			tokenString := coockie.Value

			_, err = jwt_utils.ParseAndVerify(tokenString)

			if err != nil {
				responceHandler.Error(http.StatusForbidden, "Undefined token string", "Login to continue")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
