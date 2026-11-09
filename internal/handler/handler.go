package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"ski-verleih/internal/domain"
	"ski-verleih/internal/store"
)

type Handler struct{ store *store.Store }

func New(s *store.Store) *Handler { return &Handler{store: s} }
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func decode(r *http.Request, v any) error { return json.NewDecoder(r.Body).Decode(v) }
func idFromPath(path, prefix string) (int, error) {
	return strconv.Atoi(strings.Trim(strings.TrimPrefix(path, prefix), "/"))
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Methode nicht erlaubt", 405)
		return
	}
	writeJSON(w, 200, h.store.Dashboard())
}
func (h *Handler) Skis(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, h.store.Skis())
	case http.MethodPost:
		var v domain.Ski
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		created, err := h.store.AddSki(v)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		writeJSON(w, 201, created)
	default:
		http.Error(w, "Methode nicht erlaubt", 405)
	}
}
func (h *Handler) SkiByID(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r.URL.Path, "/api/skis/")
	if err != nil {
		http.Error(w, "Ungültige ID", 400)
		return
	}
	if r.Method == http.MethodDelete {
		if h.store.DeleteSki(id) != nil {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(204)
		return
	}
	if r.Method == http.MethodPut {
		var v domain.Ski
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		if h.store.UpdateSki(id, v) != nil {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(204)
		return
	}
	http.Error(w, "Methode nicht erlaubt", 405)
}
func (h *Handler) Persons(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, h.store.Persons())
	case http.MethodPost:
		var v domain.Person
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		writeJSON(w, 201, h.store.AddPerson(v))
	default:
		http.Error(w, "Methode nicht erlaubt", 405)
	}
}
func (h *Handler) PersonByID(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r.URL.Path, "/api/persons/")
	if err != nil {
		http.Error(w, "Ungültige ID", 400)
		return
	}
	if r.Method == http.MethodDelete {
		if h.store.DeletePerson(id) != nil {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(204)
		return
	}
	if r.Method == http.MethodPut {
		var v domain.Person
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		if h.store.UpdatePerson(id, v) != nil {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(204)
		return
	}
	http.Error(w, "Methode nicht erlaubt", 405)
}
func (h *Handler) Loans(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, h.store.Loans())
	case http.MethodPost:
		var v domain.Loan
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		created, err := h.store.AddLoan(v)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		writeJSON(w, 201, created)
	default:
		http.Error(w, "Methode nicht erlaubt", 405)
	}
}
func (h *Handler) Reservations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, 200, h.store.Reservations())
	case http.MethodPost:
		var v domain.Reservation
		if decode(r, &v) != nil {
			http.Error(w, "Ungültige Daten", 400)
			return
		}
		created, err := h.store.AddReservation(v)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		writeJSON(w, 201, created)
	default:
		http.Error(w, "Methode nicht erlaubt", 405)
	}
}
