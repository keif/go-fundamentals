package main

import (
	"fmt"
	"time"
)

// lowercase because it'll only be used in this package
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// type user POINTER
// now we attach it to the struct via the initial(u user)
// moving it close to the struct isn't necessary, but… organize
func (u user) outputUserDetails() {
	// technically correct, we'd need to dereference
	// (*u).firstName, ...
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	// we don't need the keys, but it does make mapping more clear
	// appUser = user{} // null value for the struct
	// appUser = user{ ...some but not all data } // the non-included keys will be instantiated with the null values
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	// passes the pointer address
	//outputUserDetails(&appUser)
	// since we moved it as part of the struct, it doesn't need an argument passed
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
