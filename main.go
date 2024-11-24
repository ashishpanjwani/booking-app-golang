package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const eventTickets = 100

var eventName = "RhythmRiot 2024"
var remainingTickets uint = 100
var bookings = []User{}

// Alternative way to create a slice of maps -
// var bookings = make([]map[string]string, 0)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	// for loop with specific condition -
	// for remainingTickets > 0 && len(bookings) < 50
	//for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email, eventName)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			fmt.Printf("The first names of the bookings are %v\n", getFirstNames())

			if remainingTickets == 0 {
				fmt.Println("We're fully booked for this event. Please check back next year.")
				//break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or the last name you entered is incorrect")
			}
			if !isValidEmail {
				fmt.Println("Entered email address is incorrect")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets left, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
		}

		wg.Wait()

	//}

}

func greetUsers() {
	fmt.Printf("🎉 Welcome to %v! 🎉\n", eventName)
	fmt.Printf("We have a total of %v tickets available, with %v tickets still up for grabs.\n", eventTickets, remainingTickets)
	fmt.Println("🎟️   Don't miss out! Secure your tickets now and join the fun! 🎟️")
}

func getFirstNames() []string {
	// Range iterates over elements	for different data structures (so not only arrays and slices)
	// For arrays and slices, range provides the index and value for each element
	firstNames := []string{}
	// Underscore in the for loop acts as a blank identifier since we don't need the index at the moment
	// which for loop provides, so underscore will help me ignore the index variable over here
	for _, booking := range bookings {
		// Splits the string with shite space as separator
		// And returns a slice with split elements
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string, eventName string) {
	remainingTickets = remainingTickets - userTickets

	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//Create a map to store the user data
	// var user = make(map[string]string)
	// user["firstName"] = firstName
	// user["lastName"] = lastName
	// user["email"] = email
	// user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("------------------")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("------------------")
	wg.Done()
}
