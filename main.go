package main

import (
	"log"
	"net/http"

	"ski-verleih/internal/handler"
	"ski-verleih/internal/store"
)

func main() {
	st := store.New()
	h := handler.New(st)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/dashboard", h.Dashboard)
	mux.HandleFunc("/api/skis", h.Skis)
	mux.HandleFunc("/api/skis/", h.SkiByID)
	mux.HandleFunc("/api/persons", h.Persons)
	mux.HandleFunc("/api/persons/", h.PersonByID)
	mux.HandleFunc("/api/loans", h.Loans)
	mux.HandleFunc("/api/reservations", h.Reservations)
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	log.Println("Ski-Verleih läuft auf http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(mux)))
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
