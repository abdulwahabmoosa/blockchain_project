package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

type UserKey int

var userKey UserKey

func dbGetUserById(id int) (User, error) {
	user := User{
		Id:           id,
		Username:     "test",
		PasswordHash: "test",
	}

	return user, nil
}

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		user, err := dbGetUserById(id)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(userKey).(User)
	if !ok {
		status := http.StatusUnprocessableEntity
		http.Error(w, http.StatusText(status), status)
		return
	}
	render.JSON(w, r, user)
}
