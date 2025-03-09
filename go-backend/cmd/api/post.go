package main

import (
	"go-backend/internal/store"
	"net/http"
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
