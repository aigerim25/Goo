package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		endpoint := r.URL.Path
		log.Printf("%s %s %s endpoint = %s",
			time.Now().Format(time.RFC3339),
			r.Method,
			r.URL.RequestURI(),
			endpoint,
		)
		next.ServeHTTP(w, r)
	})
}
