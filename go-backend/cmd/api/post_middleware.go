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

var ErrUnauthorizedAccess = errors.New("Unauthorized Action")

// authorization middleware
func (app *application) checkPostOwnership(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromContext(r.Context())
		post := postFromContext(r.Context())

		// user can modify or delete his own post
		if post.UserID == user.ID {
			next.ServeHTTP(w, r)
			return
		}

		// check role precedence
		allowed, err := app.checkRolePrecedence(r.Context(), user, requiredRole)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		if !allowed {
			app.unauthorized(w, r, ErrUnauthorizedAccess)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) checkRolePrecedence(ctx context.Context, user *store.User, requiredRole string) (bool, error) {
	role, err := app.store.Roles.GetByName(ctx, requiredRole)
	if err != nil {
		return false, err
	}

	return role.Level <= user.Role.Level, nil
}

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
