package main

import (
	"backend/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/user", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Use(auth.UserCtx)
			r.Get("/", auth.GetUser)
			// r.Put("/", UpdateUser)
			// r.Delete("/", DeleteUser)
		})
	})

	r.Post("/login", auth.Login)

	http.ListenAndServe(":3000", r)
}
