package main

import (
	"app/routes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// Grab and cache the templates at initialisation
var cwd, _ = os.Getwd()
var pageTemplates = template.Must(template.ParseFiles(filepath.Join(cwd, "./templates/test.html")))

// Initialises and start the webapp
func main() {
	fmt.Println("Initialising the app")

	routes.RegisterRoutes()

	fmt.Println("Completed initialising the app")
	fmt.Println("App starting")

	http.ListenAndServe(":8080", nil)

	fmt.Println("App finished")
}
