package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"warranty.com/controllers"
)

func RouterInit() http.Handler {
	controllers.ControllersInit()

	r := chi.NewRouter()
	r.Use(cors.Default().Handler)

	r.Route("/token", TokenRoutes)
	r.Route("/brand", UserRoutes)

	return r
}
