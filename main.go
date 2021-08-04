package main

import (
	"fmt"

	"github.com/mrcoggsworth/pluralsight-webservice/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Chris",
		LastName:  "Scogin",
	}

	fmt.Println(u)
	name := u.PrintName()

	fmt.Println(name)
	fmt.Println(u)

	u.ChangeName("Cody", "Randell")
	newName := u.PrintName()

	fmt.Println(newName)
	fmt.Println(u)

}
