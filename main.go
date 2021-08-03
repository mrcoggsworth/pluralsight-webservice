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

}
