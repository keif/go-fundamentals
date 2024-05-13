package main

import (
	"errors"
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

// because we are modifying the values, we need to utilize the pointer
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name, and birthdate are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
