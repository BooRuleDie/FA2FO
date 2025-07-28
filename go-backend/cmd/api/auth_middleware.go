package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"go-backend/internal/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
)

var (
	ErrMissing      = errors.New("authorization header is missing")
	ErrMalformed    = errors.New("authorization header is malformed")
	ErrDecoding     = errors.New("base64 decoding failed")
	ErrInvalidCreds = errors.New("invalid credentials")
	ErrClaims       = errors.New("failed to retrieve claims of the token")
	ErrUserID       = errors.New("failed to retrieve userID from token")
)

const userCtxKey = contextKey("user")

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.unauthorizedWithError(w, r, ErrMissing)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedWithError(w, r, ErrMalformed)
			return
		}

		token := parts[1]
		jwtToken, err := app.auth.ValidateToken(token)
		if err != nil {
			app.unauthorizedWithError(w, r, ErrMalformed)
			return
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			app.unauthorizedWithError(w, r, ErrClaims)
			return
		}

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			app.unauthorizedWithError(w, r, ErrUserID)
			return
		}

		ctx := r.Context()
		user, err := app.getUser(ctx, userID)
		if err != nil {
			app.unauthorizedWithError(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {
	// Try to get the user from the cache by their userID.
	user, err := app.cache.Users.Get(ctx, userID)
	if err != nil {
		// Return error if the cache call fails completely.
		return nil, err
	}
	if user != nil {
		// If user is found in cache, return it immediately.
		// app.logger.Info("getUser from cache!")
		return user, nil
	} else {
		// app.logger.Info("getUser from db!")
		// User not found in cache, fetch from persistent store.
		user, err = app.store.Users.GetByID(ctx, userID)
		if err != nil {
			// Return error if user could not be retrieved from the database.
			return nil, err
		}
		// Cache the retrieved user for future requests.
		err = app.cache.Users.Set(ctx, user)
		if err != nil {
			// Return error if caching attempt fails.
			return nil, err
		}
	}

	// Return the retrieved user (either from cache or database).
	return user, err
}

// Get post from context
func getUserFromContext(ctx context.Context) *store.User {
	user, ok := ctx.Value(userCtxKey).(*store.User)
	if !ok {
		return nil
	}
	return user
}

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.basicAuthUnauthorized(w, r, ErrMissing)
				return
			}

			// parse it and the get the base64 encoded creds
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Basic" {
				app.basicAuthUnauthorized(w, r, ErrMalformed)
				return
			}

			// decode it
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				app.basicAuthUnauthorized(w, r, ErrDecoding)
				return
			}

			// check if the creds are valid
			username := app.config.auth.basic.username
			pass := app.config.auth.basic.pass

			creds := strings.SplitN(string(decoded), ":", 2)
			if len(creds) != 2 || creds[0] != username || creds[1] != pass {
				app.basicAuthUnauthorized(w, r, ErrInvalidCreds)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
