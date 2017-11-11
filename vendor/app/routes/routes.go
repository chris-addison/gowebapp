package routes

import (
	"app/models"
	"app/views"
	"net/http"
)

// RegisterRoutes registers all of the routes used by the app
func RegisterRoutes() {
	http.HandleFunc("/", testViewHandler)
}

func testViewHandler(responseWriter http.ResponseWriter, request *http.Request) {
	//result, _ := http.Get("https://unity3d.com//showcase/gallery/more/Default/featured/weight/1000")
	//bytes, _ := ioutil.ReadAll(result.Body)
	//fmt.Fprint(responseWriter, string(bytes))

	data := models.GetNextGame()
	error := views.Display(responseWriter, "test", data)
	if error != nil {
		http.Error(responseWriter, error.Error(), http.StatusInternalServerError)
	}
}
