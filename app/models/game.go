package models

import "sync"

// Game is a model for stroing the information for an individual title
type Game struct {
	Title       string
	Description string
	Platforms   string
}

// Storage for the singleton
var once sync.Once
var games = []Game{}

// GetNext grabs the next Game and returns it
func GetNextGame() Game {
	// TODO: lookup user and check which titles they've viewed
	return (*getTitles())[0]
}

// getTitles returns the array of Games
func getTitles() *[]Game {
	// Thread-safe lazy singleton implementation
	once.Do(func() {
		// TODO: get real data
		games = []Game{
			Game{
				"hello world",
				"This is a cool game",
				"Linux, Mac, Windows",
			},
		}
	})
	return &games
}
