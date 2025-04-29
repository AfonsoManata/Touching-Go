package main 

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName = "Go conference" // Value can change
	const conferenceTickets = 50 // Value not change 
	var remainingTickets uint = 50 // This value will change but will be positive

	fmt.Printf("conferenceTickets is %T \n", conferenceTickets) // %T type

	fmt.Printf("Welcome to our %v booking app\n", conferenceName)		
	fmt.Printf("We have total of %v and %v remaining Tickets!\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here\n")

	// In go if you assign a value you dont need the data type 

	for {

	var bookings []string // You cant mix the types	and if you dont say the size is a slice
	var firstName string
	var lastName string
	var email string
	var userTickets uint 

	fmt.Println("Enter you first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter you last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings,firstName + " " + lastName)

	fmt.Printf("Number of bookings is %v\n", len(bookings))


	// Se nao for Printf % nao funciona
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a"+ 
		"confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v!\n", remainingTickets, conferenceName)
	
	firstNames := []string{}

	for _, booking := range bookings { // _ -> unused variables
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("The first names bookings are %v\n", firstNames)

	}
} 
