package main

import (
	"context"
	"errors"
	"go-backend/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type contextKey string

const postContextKey = contextKey("post")

// getPost middleware
func (app *application) postCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pre-processing
		stringPostID := chi.URLParam(r, "postID")
		postID, err := strconv.ParseInt(stringPostID, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		post, err := app.store.Posts.GetByID(r.Context(), postID)
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

		// Add post to context
		ctx := contextWithPost(r.Context(), post)

		// Call next handler with modified context
		next.ServeHTTP(w, r.WithContext(ctx))

		// Any post-processing could go here
	})
}

// Get post from context
func postFromContext(ctx context.Context) *store.Post {
	post, ok := ctx.Value(postContextKey).(*store.Post)
	if !ok {
		return nil
	}
	return post
}

// Set post to context
func contextWithPost(ctx context.Context, post *store.Post) context.Context {
	return context.WithValue(ctx, postContextKey, post)
}
