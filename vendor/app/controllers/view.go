package controllers

import (
	"app/models"
	"lib/session"
	"lib/view"
	"net/http"
)

// ViewHandler is a controller for showing the basic view. It grabs the next game and displays it.
func ViewHandler(responseWriter http.ResponseWriter, request *http.Request, session *session.Session) {
	data, err := models.GetNextGame(session)
	if err != nil {
		http.Error(responseWriter, "Unabled to fetch next game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Display the view template with the next game
	if err := view.Display(responseWriter, "view", data); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
