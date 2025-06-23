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

const userCtx = contextKey("user")

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

		user, err := app.store.Users.GetByID(ctx, userID)
		if err != nil {
			app.unauthorizedWithError(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Get post from context
func getUserFromContext(ctx context.Context) *store.User {
	user, ok := ctx.Value(userCtx).(*store.User)
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
