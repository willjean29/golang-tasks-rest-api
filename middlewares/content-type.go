package middlewares

import (
	"log"
	"net/http"
	"strings"
)

func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		if !strings.Contains(r.URL.Path, "/uploads/") && r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
