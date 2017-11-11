package models

import "lib/session"

// Game is a model for stroing the information for an individual title
type Game struct {
	Title       string
	Description string
	Platforms   string
}

// Key used to hold the user's viewed games in the session
const vgSessionKey = "VIEWED_GAMES"

// Holds an array of Games which store info to be printed
var games = []Game{
	Game{
		"A game",
		"This is a cool game",
		"Linux, Mac, Windows",
	},
	Game{
		"Other game",
		"This is a fun game",
		"Linux, Android",
	},
}

// GetNextGame grabs the next unviewed Game for the given user and returns it
func GetNextGame(session *session.Session) Game {
	viewed := getViewedGames(session)
	// Interate through the array and return the first unviewed game found
	for _, game := range games {
		if _, exists := (*viewed)[game.Title]; !exists {
			(*viewed)[game.Title] = true
			return game
		}
	}
	// No unviewed games, so start from scratch
	*viewed = make(map[string]bool)
	(*viewed)[games[0].Title] = true
	return games[0]
}

// getViewedGames returns a "set" of games that have already been viewd by the given user
func getViewedGames(session *session.Session) *map[string]bool {
	// If session already has a "set" return a pointer to it
	if viewed, ok := session.Read(vgSessionKey).(*map[string]bool); ok && viewed != nil {
		return viewed
	}
	// Otherwise, create a new "set"
	viewed := make(map[string]bool)
	session.Write(vgSessionKey, &viewed)
	return &viewed
}
