package main

import (
	"errors"
	"go-backend/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"` // optional field, default null
}

type CreatePostResponse struct {
	ID int64 `json:"id"`
}

// CreatePost godoc
//
//	@Summary		Creates a post
//	@Description	Creates a post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreatePostPayload	true	"Post payload"
//	@Success		201		{object}	CreatePostResponse	"Contains the ID of the created post"
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts [post]
func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var post CreatePostPayload
	if err := readJSON(w, r, &post); err != nil {
		// app.jsonResponseError(w, http.StatusBadRequest, err.Error())
		app.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(post); err != nil {
		app.badRequest(w, r, err)
		return
	}

	newPost := &store.Post{
		Title:   post.Title,
		Content: post.Content,
		Tags:    post.Tags,
		UserID:  1,
	}

	if err := app.store.Posts.Create(r.Context(), newPost); err != nil {
		// app.jsonResponseError(w, http.StatusInternalServerError, err.Error())
		app.internalServerError(w, r, err)
		return
	}

	res := CreatePostResponse{ID: newPost.ID}
	if err := app.jsonResponse(w, http.StatusCreated, res); err != nil {
		// app.jsonResponseError(w, http.StatusInternalServerError, err.Error())
		app.internalServerError(w, r, err)
		return
	}
}

// GetPost godoc
//
//	@Summary		Fetches a post
//	@Description	Fetches a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Post ID"
//	@Success		200	{object}	store.Post
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{id} [get]
func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := postFromContext(r.Context())
	if post == nil {
		app.internalServerError(w, r, errors.New("failed to fetch post from middleware"))
		return
	}

	// get comments
	comments, err := app.store.Comments.GetByPostID(r.Context(), post.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	post.Comments = comments

	if err = app.jsonResponse(w, http.StatusOK, post); err != nil {
		// app.jsonResponseError(w, http.StatusInternalServerError, err.Error())
		app.internalServerError(w, r, err)
		return
	}
}

type UpdatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags" validate:"required"`
}

// UpdatePost godoc
//
//	@Summary		Updates a post
//	@Description	Updates a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Post ID"
//	@Param			payload	body		UpdatePostPayload	true	"Post payload"
//	@Success		200		{object}	store.Post
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{id} [patch]
func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	postID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// updatePostPayload
	var upp UpdatePostPayload
	if err := readJSON(w, r, &upp); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := Validate.Struct(upp); err != nil {
		app.badRequest(w, r, err)
		return
	}
	
	user := getUserFromContext(r.Context())
	if user == nil {
		app.internalServerError(w, r, errors.New("nil user struct retrieved from getUserFromContext"))
		return
	}

	updatedPost := &store.Post{
		ID:      postID,
		Title:   upp.Title,
		Content: upp.Content,
		Tags:    upp.Tags,
		UserID: user.ID,
	}

	err = app.store.Posts.Update(r.Context(), updatedPost)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFound(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

// DeletePost godoc
//
//	@Summary		Deletes a post
//	@Description	Delete a post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Post ID"
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{id} [delete]
func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r.Context())
	if user == nil {
		app.internalServerError(w, r, ErrNilUser)
		return
	}
	
	idParam := chi.URLParam(r, "postID")
	postID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	deletedPost := &store.Post{
		ID: postID,
		UserID: user.ID,
	}

	err = app.store.Posts.Delete(r.Context(), deletedPost)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFound(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}


