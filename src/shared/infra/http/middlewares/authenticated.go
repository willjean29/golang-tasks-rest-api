package middlewares

import (
	token "app/src/shared/adapters/token"
	error "app/src/shared/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

var tokenAdapter token.TokenAdapter = token.NewJwtAdapter()

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(error.New("Unauthorized", http.StatusUnauthorized, nil))
			return
		}
		token := cookie.Value
		isValidToken, tokenDecoded, err := tokenAdapter.ValidateToken(token)
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
