package routes

import (
	"github.com/go-chi/chi"
	"warranty.com/controllers"
	"warranty.com/middleware"
)

func TokenRoutes(r chi.Router) {
	r.Get("/nonce/{productId}", controllers.GetTokenNonce)
	r.Post("/token", controllers.GetToken)
	r.Post("/register", controllers.RegisterToken)
	r.Post("/owner", controllers.GetTokensByOwner)
	r.Group(func(r chi.Router) {
		r.Use(middleware.VerifyAddress)
		r.Post("/all", controllers.GetTokensByBrand)
		r.Post("/approve", controllers.ApproveToken)
		r.Post("/add", controllers.AddApprovedToken)
		r.Post("/claim", controllers.SetClaim)
	})
}
