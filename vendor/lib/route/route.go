package route

import (
	"lib/session"
	"net/http"
)

// CreateRoute initialises a http request handler. This is a wrapper for the http.HandleFunc function
// the main difference is that create route provides the current session. The session is locked for the duration of the request
// to avoid any concurrency issues. Locking sessions in use is a common practice.
func CreateRoute(pattern string, function func(responseWriter http.ResponseWriter, request *http.Request, session *session.Session)) {
	viewHandler := func(responseWriter http.ResponseWriter, request *http.Request) {
		session := session.GetManager().Start(responseWriter, request)
		session.Lock.Lock()
		defer session.Lock.Unlock()
		function(responseWriter, request, session)
	}
	http.HandleFunc(pattern, viewHandler)
}
