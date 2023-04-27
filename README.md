<img src="https://cdn.sriggle.tech/kantents/production/1/1345/09/6a505ba0-0302-4df7-a89d-e963a579c172.webp" alt="Sample Image" width="400" height="300">


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
* The reservation system should only allow valid inputs for the number of seats and seat numbers.
* The reservation_test.go file should not be edited.
* You should only edit the reservation.go file in the ticket package.
* Your implementation should use channels and goroutines to implement concurrency in the reservation system.
* You should not use any external libraries other than the Go standard library.
* The code should compile with the most recent version of the Go compiler.
* The program should not panic under any circumstances.

**Motivation**

* Concurrency is an increasingly important concept in modern programming, as more and more applications are designed to take advantage of multiple processors and distributed systems. Go's built-in support for concurrency makes it a popular choice for building highly-scalable applications that can handle large numbers of users and requests. This assignment will help students understand how to use goroutines and channels to build concurrent applications in Go.


**Submission**

* Commit and push your working code to your GIT repository.
* Ensure all tests pass otherwise you will receive no credit.
