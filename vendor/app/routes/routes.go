package routes

import (
	"app/controllers"
	"lib/route"
)

// RegisterRoutes registers all of the routes used by the app
func RegisterRoutes() {
	route.CreateRoute("/view", controllers.ViewHandler)
}
