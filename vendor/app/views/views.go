package views

import (
	"html/template"
	"net/http"
	"sync"
)

// Values to determine how and where templates are stored
const path = "./templates/"
const ext = ".html"

var (
	// Lockable map for storing cached views
	// Number of users will be low, so no need for a RW lock here
	viewMapLock sync.Mutex
	viewMap     = make(map[string]*template.Template)
)

// Get the template thet corresponds to the given name
// Lazily loads the templates
func getTemplate(name string) *template.Template {
	// Unlock the map and defer the unlock for when the function returns
	// Defer is needed as template.Must will panic if it fails to parse a template
	viewMapLock.Lock()
	defer viewMapLock.Unlock()

	// Create or load the template
	templ, exists := viewMap[name]
	if !exists {
		templ = template.Must(template.ParseFiles(path+name+ext, path+"base"+ext))
		viewMap[name] = templ
	}
	return templ
}

// Display is a function that when given a ResponseWriter, a template name, and data will display
// the correspondng template on the page
func Display(responseWriter http.ResponseWriter, templateName string, data interface{}) error {
	return getTemplate(templateName).ExecuteTemplate(responseWriter, "base", data)
}
