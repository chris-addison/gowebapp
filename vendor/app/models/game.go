package models

import (
	"errors"
	"fmt"
	"lib/session"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Game is a model for stroing the information for an individual title
type Game struct {
	Title       string
	Developer   string
	Description string
	Genre       string
	Platforms   string
	Image       string
}

// Key used to hold the user's viewed games in the session
const vgSessionKey = "VIEWED_GAMES"

// URL to get the game list from
const gameURL = "https://unity3d.com//showcase/gallery/more/Default/featured/weight/1000"

// The frequency the games list is refreshed
const fetchFrequency = 30 * time.Minute

// Holds an array of Games which store info to be printed
var games = []Game{}

// Initialise the games list
func init() {
	// Load games list on start
	getGamesList()

	// Asycnhronous function to update the games list periodically
	go func() {
		for {
			<-time.After(fetchFrequency)
			getGamesList()
		}
	}()
}

// GetNextGame grabs the next unviewed Game for the given user and returns it
func GetNextGame(session *session.Session) (Game, error) {
	// Catch invalid state of the array of games
	if len(games) < 1 {
		return Game{}, errors.New("no games found")
	}

	viewed := getViewedGames(session)
	// Interate through the array and return the first unviewed game found
	for _, game := range games {
		if _, exists := (*viewed)[game.Title]; !exists {
			(*viewed)[game.Title] = true
			return game, nil
		}
	}
	// No unviewed games, so start from scratch
	*viewed = make(map[string]bool)
	(*viewed)[games[0].Title] = true
	return games[0], nil
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

// getGamesList updates the game list from url in gameURL
func getGamesList() {
	doc, err := goquery.NewDocument(gameURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Walk the html and scrape the game details
	gameSelectors := doc.Find("li.game")
	temp := make([]Game, gameSelectors.Length())
	gameSelectors.Each(func(i int, selection *goquery.Selection) {
		game := Game{}
		game.Title = selection.Find(".title").Text()
		game.Developer = selection.Find(".developer").Text()
		game.Description = selection.Find(".description").Text()
		game.Genre = selection.Find(".genres").Text()
		game.Platforms = strings.Join(selection.Find(".tip").Map(func(i int, platSelect *goquery.Selection) string {
			return platSelect.Text()
		}), ", ")
		game.Image = selection.Find(".ic").AttrOr("src", "")
		temp[i] = game
	})

	// Update game list if successful
	if len(temp) > 0 {
		games = temp
	}
}
