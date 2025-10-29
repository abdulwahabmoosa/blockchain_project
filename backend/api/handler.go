package api

import (
	"backend/auth"
	"backend/db"
	"backend/db/models"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type RequestHandler struct {
	db *db.Database
}

func NewRequestHandler(db *db.Database) *RequestHandler {
	return &RequestHandler{db}
}

func (handler *RequestHandler) Start() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Route("/{id}", func(r chi.Router) {
			r.Use(handler.UserCtx)
			r.Get("/", GetUser)
			// r.Put("/", UpdateUser)
			// r.Delete("/", DeleteUser)
		})
	})

	r.Post("/login", auth.Login)

	http.ListenAndServe(":3000", r)
}

func (handler *RequestHandler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		user, err := handler.db.GetUserById(id)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "users", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, ok := ctx.Value("users").(models.User)
	if !ok {
		status := http.StatusUnprocessableEntity
		http.Error(w, http.StatusText(status), status)
		return
	}

	render.JSON(w, r, user)
}
