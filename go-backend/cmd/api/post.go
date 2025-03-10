package main

import (
	"errors"
	"go-backend/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type CreatePostResponse struct {
	ID int64 `json:"id"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var post CreatePostPayload
	if err := readJSON(w, r, &post); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	newPost := &store.Post{
		Title:   post.Title,
		Content: post.Content,
		Tags:    post.Tags,
		UserID:  1,
	}

	if err := app.store.Posts.Create(r.Context(), newPost); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := CreatePostResponse{ID: newPost.ID}
	if err := writeJSON(w, http.StatusCreated, res); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	postID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		// you probably don't want to return those
		// errors in production, as it can give so much
		// clue to the hackers
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := app.store.Posts.GetByID(r.Context(), postID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			writeJSONError(w, http.StatusNotFound, store.ErrNotFound.Error())
			return
		default:
			writeJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err = writeJSON(w, http.StatusOK, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
