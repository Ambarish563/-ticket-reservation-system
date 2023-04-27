package tickets

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
)

type Reservation struct {
	mtx   sync.Mutex
	seats []bool
}

// NewReservation(numSeats int) *Reservation: a constructor function that creates a new Reservation\n
// object with a seats field of length numSeats (all initially set to false).

func NewReservation(numSeats int) *Reservation {
	return &Reservation{seats: make([]bool, numSeats)}
}

// CheckAvailability(numSeats int) bool: a method that checks if numSeats seats are available for reservation.
// It returns true if available, false otherwise.
func (r *Reservation) CheckAvailability(numSeats int) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if numSeats > len(r.seats) {
		return false
	}
	for i := 0; i < len(r.seats); i++ {
		if !r.seats[i] {
			available := true
			for j := 1; j < numSeats && i+j < len(r.seats); j++ {
				if r.seats[i+j] {
					available = false
					i += j
					break
				}
			}
			if available {
				return true
			}
		}
	}
	return false
}

// ReserveSeats(numSeats int) ([]int, error): a method that reserves numSeats seats if available,\n
// and returns their seat numbers as a slice of integers. If the seats are not available, it returns nil and an error.
func (r *Reservation) ReserveSeats(numSeats int) ([]int, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if numSeats > len(r.seats) {
		return nil, errors.New("not enough seats available")
	}
	var seats []int
	for i := 0; i <= len(r.seats)-numSeats; i++ {
		if !r.seats[i] {
			available := true
			for j := 1; j < numSeats; j++ {
				if r.seats[i+j] {
					available = false
					i += j
					break
				}
			}
			if available {
				for j := 0; j < numSeats; j++ {
					r.seats[i+j] = true
					seats = append(seats, i+j+1)
				}
				return seats, nil
			}
		}
	}
	return nil, errors.New("not enough seats available")
}

// CancelReservation is a method of the Reservation struct. It takes a slice of seat numbers as input and returns an error.
// It cancels the reservation of the seats in the slice by setting their corresponding values in the seats slice to false.
// If any seat number in the slice is invalid or not reserved, the method returns an error.
func (r *Reservation) CancelReservation(seats []int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	for _, seat := range seats {
		if seat < 1 || seat > len(r.seats) {
			return fmt.Errorf("invalid seat number: %d", seat)
		}
		if !r.seats[seat-1] {
			return fmt.Errorf("seat %d is not reserved", seat)
		}
		r.seats[seat-1] = false
	}
	return nil
}

type Confirmation struct {
	seats []int
}

// NewConfirmation(seats []int) *Confirmation: a constructor function that creates a new Confirmation object with a seats field set to seats.
func NewConfirmation(seats []int) *Confirmation {
	return &Confirmation{seats: seats}
}

// GenerateConfirmationCode() string: a method that generates a random confirmation code and returns it as a string.
func (c *Confirmation) GenerateConfirmationCode() string {
	codeBytes := make([]byte, 4)
	if _, err := rand.Read(codeBytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(codeBytes)
}

// PrintConfirmationDetails(): a method that prints the details of the reservation, including the seats reserved and the confirmation code.
func (c *Confirmation) PrintConfirmationDetails() {
	fmt.Printf("Seats reserved: %v\nConfirmation code: %s\n", c.seats, c.GenerateConfirmationCode())
}
