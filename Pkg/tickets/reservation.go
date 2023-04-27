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

	// Create a channel to communicate the result of seat availability checks.
	available := make(chan bool)

	// Start a goroutine to check the availability of each seat in a separate thread.
	for i := 0; i < len(r.seats); i++ {
		go func(i int) {
			if !r.seats[i] {
				// Check the availability of numSeats seats starting from seat i.
				// If any of the seats is already reserved, mark the range as unavailable.
				for j := 1; j < numSeats && i+j < len(r.seats); j++ {
					if r.seats[i+j] {
						available <- false
						return
					}
				}
				// If all numSeats seats starting from seat i are available, mark the range as available.
				available <- true
			} else {
				// If seat i is already reserved, mark the range as unavailable.
				available <- false
			}
		}(i)
	}

	// Wait for all goroutines to finish and collect the results of seat availability checks.
	for i := 0; i < len(r.seats); i++ {
		if <-available {
			return true
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
	type seatRange struct {
		start int
		end   int
	}
	seatRanges := make(chan seatRange)
	go func() {
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
					seatRanges <- seatRange{i, i + numSeats - 1}
				}
			}
		}
		close(seatRanges)
	}()
	var seats []int
	for sr := range seatRanges {
		for j := sr.start; j <= sr.end; j++ {
			r.seats[j] = true
			seats = append(seats, j+1)
		}
		return seats, nil
	}
	return nil, errors.New("not enough seats available")
}

// CancelReservation is a method of the Reservation struct. It takes a slice of seat numbers as input and returns an error.
func (r *Reservation) CancelReservation(seats []int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	// Create a channel to receive errors from the goroutines.
	errChan := make(chan error)

	// Create a waitgroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Iterate through the seats and create a goroutine for each seat.
	for _, seat := range seats {
		wg.Add(1)
		go func(seat int) {
			defer wg.Done()

			// Check if the seat number is valid and the seat is reserved.
			if seat < 1 || seat > len(r.seats) {
				errChan <- fmt.Errorf("invalid seat number: %d", seat)
				return
			}
			if !r.seats[seat-1] {
				errChan <- fmt.Errorf("seat %d is not reserved", seat)
				return
			}

			// Cancel the reservation for the seat.
			r.seats[seat-1] = false
		}(seat)
	}

	// Close the channel after all goroutines have finished.
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Check for errors returned from the goroutines.
	for err := range errChan {
		if err != nil {
			return err
		}
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
