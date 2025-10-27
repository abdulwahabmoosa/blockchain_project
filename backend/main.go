package main

import (
	"context"
	"net/http"

	"backend/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func dbGetUser(username string) (*auth.User, error) {
	return &auth.User{
			Username: username,
			Email:    username + "@email.com",
			Password: "test"},
		nil
}

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		user, err := dbGetUser(username)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value("user").(*auth.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	render.JSON(w, r, user)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Route("/{username}", func(r chi.Router) {
			r.Use(UserCtx)
			r.Get("/", GetUser)
			// r.Put("/", UpdateUser)
			// r.Delete("/", DeleteUser)
		})
	})

	http.ListenAndServe(":3000", r)
}
