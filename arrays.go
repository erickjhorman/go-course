package main

import "fmt"

func main() {
	var remainingTickets = 50

	var booking = [50]string{"Maria", "Erick", "Romero"} //declare an array
	var booking2 = [50]string{}
	//var array [50]string

	//adding elements
	booking2[0] = "Nana"
	booking[1] = "Erick"

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
	booking[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", booking)
	fmt.Printf("The first name: %v\n", booking[0])
	fmt.Printf("Array type: %T\n", booking)
	fmt.Printf("Array lenght: %v\n", len(booking))
}
