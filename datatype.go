package main

import "fmt"

func main() {

	//datatype
	var userName string
	var userTickets int
	userName = "Tom"
	userTickets = 1
	var positiveNumber uint = 10

	syntacticSugar := "Syntactic Sugar" //Syntactic Sugar
	fmt.Print(positiveNumber)
	fmt.Print(syntacticSugar)
	fmt.Print(userName)
	fmt.Print(userTickets)

	fmt.Printf("User %v booked %v tickets.\n,", userName, userTickets) //%T check variable type

	fmt.Printf("username is %T and tickets is %T",
		userName, userTickets)

}
