package routes

import (
	"app/models"
	"app/views"
	"lib/session"
	"net/http"
)

// createRoute initialises a http request handler. This is a wrapper for the http.HandleFunc function
// the main difference is that create route provides the current session. The session is locked for the duration of the request
// to avoid any concurrency issues. Locking sessions in use is a common practice.
func createRoute(pattern string, function func(responseWriter http.ResponseWriter, request *http.Request, session *session.Session)) {
	viewHandler := func(responseWriter http.ResponseWriter, request *http.Request) {
		session := session.GetManager().Start(responseWriter, request)
		session.Lock.Lock()
		defer session.Lock.Unlock()
		function(responseWriter, request, session)
	}
	http.HandleFunc(pattern, viewHandler)
}

// RegisterRoutes registers all of the routes used by the app
func RegisterRoutes() {
	createRoute("/view", testViewHandler)
}

func testViewHandler(responseWriter http.ResponseWriter, request *http.Request, session *session.Session) {
	data, err := models.GetNextGame(session)
	if err != nil {
		http.Error(responseWriter, "Unabled to fetch next game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := views.Display(responseWriter, "test", data); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
