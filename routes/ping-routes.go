package routes

import (
	"github.com/go-chi/chi"
	"warranty.com/controllers"
)

func PingRoutes(r chi.Router) {
	r.Get("/", controllers.Ping)
}
