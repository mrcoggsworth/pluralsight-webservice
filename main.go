package main

import (
	// "fmt"
	"net/http"

	"github.com/mrcoggsworth/pluralsight-webservice/controllers"
	// "github.com/mrcoggsworth/pluralsight-webservice/models"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
