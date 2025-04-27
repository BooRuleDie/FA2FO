package main

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

var (
	ErrMissing = errors.New("authorization header is missing")
	ErrMalformed = errors.New("authorization header is malformed")
	ErrDecoding = errors.New("base64 decoding failed")
	ErrInvalidCreds = errors.New("invalid credentials")
)

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
