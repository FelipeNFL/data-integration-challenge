package core

import (
	"io"
	"log"
	"bytes"
	"net/http"
)

func ParseBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        body, err := io.ReadAll(r.Body)
        
		if err != nil {
            log.Printf("Error reading body: %v", err)
            http.Error(w, "can't read body", http.StatusBadRequest)
            return
        }

        r.Body = io.NopCloser(bytes.NewBuffer(body))

        next.ServeHTTP(w, r)
	})
}

func SetDbMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := core.GetDb()

		if err != nil {
			log.Fatal(err)
		}

		ctx := context.WithValue(r.Context(), "DB", db)
		next.ServeHTTP(w, r.WithContext(ctx))
    })
}