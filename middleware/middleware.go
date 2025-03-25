package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		log.Println(time.Since(now), r.Method, r.URL.Path)
		log.Println("From", r.URL.Host)
		f(w, r)
	}
}
