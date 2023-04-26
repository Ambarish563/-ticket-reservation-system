package tickets

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReservation_CheckAvailability(t *testing.T) {
	r := NewReservation(10)
	if !r.CheckAvailability(5) {
		t.Errorf("expected true, got false")
	}
	if r.CheckAvailability(11) {
		t.Errorf("expected false, got true")
	}
}

func TestReservation_ReserveSeats(t *testing.T) {
	r := NewReservation(10)
	seats, err := r.ReserveSeats(5)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if len(seats) != 5 {
		t.Errorf("expected 5 seats reserved, got %d", len(seats))
	}
	seats, err = r.ReserveSeats(6)
	if err == nil {
		t.Errorf("expected error but we got no error")
	}
	if seats != nil {
		t.Errorf("expected nil seats, got %v", seats)
	}
}

func TestReservation_CancelReservation(t *testing.T) {
	r := NewReservation(10)
	seats, _ := r.ReserveSeats(5)
	err := r.CancelReservation(seats)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	err = r.CancelReservation([]int{11})
	if err == nil {
		t.Errorf("expected non-nil error, got nil")
	}
	err = r.CancelReservation([]int{1, 3, 5})
	if err == nil {
		t.Errorf("expected non-nil error, got nil")
	}
}

func TestConfirmation_GenerateConfirmationCode(t *testing.T) {
	c := NewConfirmation([]int{1, 2, 3, 4, 5})
	code := c.GenerateConfirmationCode()
	if code == "" {
		t.Errorf("expected non-empty code, got empty string")
	}
	if len(code) != 8 {
		t.Errorf("expected 8-character code, got %d characters", len(code))
	}
}

func TestConfirmation_PrintConfirmationDetails(t *testing.T) {
	c := NewConfirmation([]int{1, 2, 3})

	// Capture output from PrintConfirmationDetails
	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	c.PrintConfirmationDetails()

	w.Close()
	os.Stdout = oldOutput

	outputBytes, _ := ioutil.ReadAll(r)
	if len(string(outputBytes)) == 0 {
		t.Errorf("It is not Printing the Confirmation Details ")
	}
}
