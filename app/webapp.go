package webapp

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"webapp/app/models"
)

// Grab and cache the templates at initialisation
var cwd, _ = os.Getwd()
var pageTemplates = template.Must(template.ParseFiles(filepath.Join(cwd, "./templates/test.html")))

// Start initialises and starts the webapp
func Start() {
	fmt.Print("test")

	//routes.registerRoutes()

	http.HandleFunc("/", testViewHandler)

	http.ListenAndServe(":8080", nil)
}

func testViewHandler(responseWriter http.ResponseWriter, request *http.Request) {
	//result, _ := http.Get("https://unity3d.com//showcase/gallery/more/Default/featured/weight/1000")
	//bytes, _ := ioutil.ReadAll(result.Body)
	//fmt.Fprint(responseWriter, string(bytes))

	data := models.GetNextGame()
	error := pageTemplates.ExecuteTemplate(responseWriter, "test.html", data)
	if error != nil {
		http.Error(responseWriter, error.Error(), http.StatusInternalServerError)
	}
}
