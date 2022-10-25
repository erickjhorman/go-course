package main

import "fmt"

func main() {
	var remainingTickets = 50
	var testPointer uint = 8
	var firstName string
	var lastName string
	var email string
	tickets := 0

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&tickets)

	remainingTickets = remainingTickets - tickets

	/* What's a pointer
	It's a variable that points to the memory address of another variable
	*/
	fmt.Print(&testPointer)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,
		lastName, tickets, email)

}
