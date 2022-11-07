package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"warranty.com/controllers"
)

func RouterInit(version string) http.Handler {
	controllers.ControllersInit()

	r := chi.NewRouter()
	r.Use(cors.Default().Handler)

	route := fmt.Sprintf("/api/%s", version)
	r.Route(route, routes)

	return r
}

func routes(r chi.Router) {
	r.Route("/ping", PingRoutes)
	r.Route("/token", TokenRoutes)
	r.Route("/brand", UserRoutes)
}
