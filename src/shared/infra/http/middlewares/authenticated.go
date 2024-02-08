package middlewares

import (
	token "app/providers/TokenProvider"
	error "app/src/shared/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

var tokenProvider token.TokenProvider = token.NewJwtProvider()

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(error.New("Unauthorized", http.StatusUnauthorized, nil))
			return
		}
		token := cookie.Value
		isValidToken, tokenDecoded, err := tokenProvider.ValidateToken(token)
		if !isValidToken {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(error.New("Unauthorized (Token invalid)", http.StatusUnauthorized, err))
			return
		}

		id := tokenDecoded["userId"].(string)
		userId, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(error.New("Unauthorized (Error decoded token)", http.StatusUnauthorized, err))
			return
		}
		ctx := context.WithValue(r.Context(), "userId", userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
