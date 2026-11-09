package store

import (
	"testing"

	"ski-verleih/internal/domain"
)

func TestEmptyCollectionsAreJSONFriendly(t *testing.T) {
	s := New()
	if s.Reservations() == nil {
		t.Fatal("Reservations() soll einen leeren Slice statt nil liefern")
	}
}

func TestAddLoanValidatesPersonBeforeChangingSki(t *testing.T) {
	s := New()
	_, err := s.AddLoan(domain.Loan{SkiID: 1, PersonID: 999, Date: "2026-11-06"})
	if err == nil {
		t.Fatal("Ausleihe mit unbekannter Person muss fehlschlagen")
	}
	for _, ski := range s.Skis() {
		if ski.ID == 1 && ski.Status != "available" {
			t.Fatalf("Ski-Status wurde trotz fehlgeschlagener Ausleihe geändert: %s", ski.Status)
		}
	}
}

func TestCannotDeleteRentedSki(t *testing.T) {
	s := New()
	if err := s.DeleteSki(2); err == nil {
		t.Fatal("Ausgeliehener Ski darf nicht gelöscht werden")
	}
}

func TestReservationRequiresExistingPersonAndSki(t *testing.T) {
	s := New()
	if _, err := s.AddReservation(domain.Reservation{SkiID: 1, PersonID: 999, Date: "2026-11-06"}); err == nil {
		t.Fatal("Reservierung mit unbekannter Person muss fehlschlagen")
	}
	if _, err := s.AddReservation(domain.Reservation{SkiID: 999, PersonID: 1, Date: "2026-11-06"}); err == nil {
		t.Fatal("Reservierung mit unbekanntem Ski muss fehlschlagen")
	}
}

func TestLoanCalculatesFlexibleRentalPrice(t *testing.T) {
	s := New()
	loan, err := s.AddLoan(domain.Loan{SkiID: 1, PersonID: 1, Date: "2026-11-06", Days: 3})
	if err != nil {
		t.Fatalf("AddLoan() Fehler: %v", err)
	}
	if loan.DailyPrice != 25.00 || loan.TotalPrice != 75.00 {
		t.Fatalf("Falscher Mietpreis: Preis/Tag %.2f, Gesamt %.2f", loan.DailyPrice, loan.TotalPrice)
	}
}
