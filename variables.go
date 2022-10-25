package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 //this is a constant variable
	var remainingTickets = 50
	const constant = 10 //another way to declare const
	fmt.Println("Welcome to our ", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, " tickets and", remainingTickets)
	fmt.Println("Welcome to our ", conferenceName, "booking application")

	//format code
	fmt.Printf("Hey %v", "Erick")
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)

}
