package models

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func (u *User) PrintName() string {
	return fmt.Sprintf("First Name: %s, Last Name: %s", u.FirstName, u.LastName)
}

func (u *User) ChangeName(newFirstName, newLastName string) {
	u.FirstName = newFirstName
	u.LastName = newLastName

}

var (
	users  []*User
	nextID = 1
)
