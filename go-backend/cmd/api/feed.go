package main

import (
	"go-backend/internal/store"
	"net/http"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	feedPagination, err := store.FeedPaginationParse(r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(feedPagination); err != nil {
		app.badRequest(w, r, err)
		return
	}

	// TODO: update the line after auth is handled
	var userID int64 = 1

	posts, err := app.store.Posts.Feed(r.Context(), userID, feedPagination)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, posts); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
