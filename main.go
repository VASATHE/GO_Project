package main

import (
	"fmt"
	"strings"
)

type Booking struct {
	fname          string
	lname          string
	email          string
	num_of_tickets int
}

var bookings []Booking

const totalTickets = 50

var remainTickets = 50

func main() {
	greet()
	for {
		if remainTickets == 0 {
			fmt.Println("\nAll tickets are booked!")
			showbooking()
			return
		}
		var fname string
		for {
			fmt.Println("\nEnter your first name: ")
			fmt.Scan(&fname)
			if len(fname) > 2 {
				break
			}
			fmt.Println("First name must be greater than 2 characters")
		}
		var lname string
		for {
			fmt.Println("Enter your last name: ")
			fmt.Scan(&lname)
			if len(lname) > 2 {
				break
			}
			fmt.Println("Last name must be greater than 2 characters")
		}
		var email string
		for {
			fmt.Println("Enter your email: ")
			fmt.Scan(&email)
			if strings.Contains(email, "@") {
				break
			}
			fmt.Println("Email must contain @")
		}
		var num_of_tickets int
		for {
			fmt.Println("Enter the number of tickets: ")
			fmt.Scan(&num_of_tickets)
			if num_of_tickets > 0 && num_of_tickets <= remainTickets {
				break
			}
			fmt.Printf("Enter valid ticket count (remaining tickets %v)", remainTickets)
		}
		bookings = append(bookings, Booking{
			fname:          fname,
			lname:          lname,
			email:          email,
			num_of_tickets: num_of_tickets,
		})
		remainTickets = remainTickets - num_of_tickets
		fmt.Println("Succesfully bookes tickes!")
		fmt.Printf("\nRemaining tickets are : %v", remainTickets)
	}
}

func greet() {
	fmt.Println("\nWelcome to Ticket Booking System")
	fmt.Printf("\nTotal Tickets: %v", totalTickets)
	fmt.Printf("\nCurrently available: %v", remainTickets)
}
func showbooking() {
	fmt.Println("\nList of booked tickets:")
	for i, b := range bookings {
		fmt.Printf("\n %v) %v %v %v - %v tickets", i+1, b.fname, b.lname, b.email, b.num_of_tickets)
	}
}
