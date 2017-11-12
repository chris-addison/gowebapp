package main

import (
	"app/routes"
	"fmt"
	"net/http"
)

// Initialises and start the webapp
func main() {
	fmt.Println("Initialising the app")

	routes.RegisterRoutes()

	fmt.Println("Completed initialising the app")
	fmt.Println("App starting")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("resources/css"))))
	http.ListenAndServe(":8080", nil)

	fmt.Println("App finished")
}
