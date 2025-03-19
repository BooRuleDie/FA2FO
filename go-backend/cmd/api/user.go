package main

import (
	"errors"
	"go-backend/internal/store"
	"net/http"
)

var (
	ErrSelfFollow   = errors.New("users cannot follow themselves")
	ErrSelfUnfollow = errors.New("users cannot unfollow themselves")
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromContext(r.Context())
	if userID == 0 {
		app.internalServerError(w, r, ErrUserContext)
	}

	user, err := app.store.Users.GetByID(r.Context(), userID)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFound(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromContext(r.Context())
	if userID == 0 {
		app.internalServerError(w, r, ErrUserContext)
	}

	// TODO: change that line after auth is completed
	var followerID int64 = 1

	if userID == followerID {
		app.conflict(w, r, ErrSelfFollow)
		return
	}

	err := app.store.Users.Follow(r.Context(), userID, followerID)
	if err != nil {
		switch err {
		case store.ErrAlreadyFollowing:
			app.conflict(w, r, store.ErrAlreadyFollowing)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromContext(r.Context())
	if userID == 0 {
		app.internalServerError(w, r, ErrUserContext)
	}

	// TODO: change that line after auth is completed
	var followerID int64 = 1

	if userID == followerID {
		app.conflict(w, r, ErrSelfUnfollow)
		return
	}

	if err := app.store.Users.Unfollow(r.Context(), userID, followerID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
