package routes

import (
	"app/models"
	"app/views"
	"fmt"
	"lib/session"
	"net/http"
)

// RegisterRoutes registers all of the routes used by the app
func RegisterRoutes() {
	http.HandleFunc("/view", testViewHandler)
}

func testViewHandler(responseWriter http.ResponseWriter, request *http.Request) {
	//result, _ := http.Get("https://unity3d.com//showcase/gallery/more/Default/featured/weight/1000")
	//bytes, _ := ioutil.ReadAll(result.Body)
	//fmt.Fprint(responseWriter, string(bytes))

	currentSession := session.GetManager().Start(responseWriter, request)

	fmt.Println(currentSession.Read("test"))
	currentSession.Write("test", "hello")

	data := models.GetNextGame()

	if error := views.Display(responseWriter, "test", data); error != nil {
		http.Error(responseWriter, error.Error(), http.StatusInternalServerError)
	}
}
