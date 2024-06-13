package handlers

import(
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"learn-go-api/internal/middleware"

)

func Handler (r *chi.Mux){		//	capital letter function name means function is not private and is exported
	
	// Global middleware
	// applied all the time, to all endpoints we make
	r.Use(chimiddle.StripSlashes)	// remove trailing slashes at the end of every endpoint

	r.Route("/account", func (router chi.Router){

		// Middleware for /account route
		router.Use(middleware.Authorization)	//	check if authorized first, if not throw error, nothing else gets executed

		router.Get("/coins",GetCoinBalance)
	})

}

