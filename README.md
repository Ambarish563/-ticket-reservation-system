<img src="https://thumbs.dreamstime.com/z/online-tickets-service-purchase-check-plane-online-tickets-service-purchase-cartoon-flat-vector-illustration-man-woman-use-156679240.jpg" alt="Sample Image" width="400" height="300">


# Overview

* In this assignment, you will implement a ticket reservation system that uses goroutines and channels to simulate a group of customers reserving tickets from a group of ticket agents. The system will allow customers to reserve a ticket and receive a confirmation, while ticket agents will sell tickets to customers who have made a reservation. The goal of this assignment is to provide hands-on experience in using goroutines and channels in a concurrent application.

## Learning Objectives

* Understanding of goroutines and channels in Go.
* Familiarity with concurrency patterns in Go.
* Implementation of a concurrent application using goroutines and channels.

### Requirements:

**Your implementation of the ticket reservation system should meet the following requirements:**

* The Reservation type should have methods for checking seat availability, reserving seats, and canceling reservations. These methods should use channels to communicate with other goroutines and ensure thread safety.
* The Confirmation type should have methods for generating confirmation codes and printing confirmation details to the console.
* The implementation should be free of deadlock, race conditions, and livelock.
* The implementation should handle concurrent access to the reservation system and ensure that reservations are not double-booked.
* The implementation should support multiple users making reservations simultaneously.
* The reservation system should only allow valid inputs for the number of seats and seat numbers. Invalid inputs should be rejected and an error message should be displayed.
* The main.go file should not be edited and should simulate user interactions with the reservation system.
* You should only edit the reservation.go file in the ticket package.
* Your implementation should use channels and goroutines to implement concurrency in the reservation system.
* You should not use any external libraries other than the Go standard library.
* The code should compile with the most recent version of the Go compiler.
* The program should not panic under any circumstances.

**Submission**

* Commit and push your working code to your GIT repository.
* Ensure all tests pass otherwise you will receive no credit.
