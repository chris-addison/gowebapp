package routes

import (
	"app/models"
	"app/views"
	"lib/session"
	"net/http"
)

// RegisterRoutes registers all of the routes used by the app
func RegisterRoutes() {
	http.HandleFunc("/view", testViewHandler)
}

func testViewHandler(responseWriter http.ResponseWriter, request *http.Request) {
	currentSession := session.GetManager().Start(responseWriter, request)

	data := models.GetNextGame(currentSession)

	if error := views.Display(responseWriter, "test", data); error != nil {
		http.Error(responseWriter, error.Error(), http.StatusInternalServerError)
	}
}
