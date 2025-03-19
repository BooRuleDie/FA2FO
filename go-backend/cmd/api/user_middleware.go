package main

import (
	"context"
	"errors"
	"go-backend/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const userIDContextKey = contextKey("userID")

var ErrUserContext = errors.New("failed to retrieve userID from userIDFromContext")

// get userID middleware
func (app *application) userIDCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pre-processing
		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			app.badRequest(w, r, store.ErrInvalidUserID)
			return
		}

		// Add userID to context
		ctx := contextWithUserID(r.Context(), userID)

		// Call next handler with modified context
		next.ServeHTTP(w, r.WithContext(ctx))

		// Any post-processing could go here
	})
}

// Get userID from context
func userIDFromContext(ctx context.Context) int64 {
	userID, ok := ctx.Value(userIDContextKey).(int64)
	if !ok {
		return 0
	}
	return userID
}

// Set userID to context
func contextWithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDContextKey, userID)
}
