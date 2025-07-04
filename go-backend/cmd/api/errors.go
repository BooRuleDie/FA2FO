package main

import (
	"net/http"
)

// errorType: method path, error: couldn't parse int...
// internal server error: GET /api/v1/auth/login, error: couldn't parse int...
const logTemplate = "%s: %s %s, error: %s"

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf(logTemplate, "Internal Server Error", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusInternalServerError, "something went wrong")
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf(logTemplate, "Bad Request", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Infof(logTemplate, "Not Found", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) conflict(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf(logTemplate, "Conflict", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf(logTemplate, "Unauthorized", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedWithError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf(logTemplate, "Unauthorized", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusUnauthorized, err.Error())
}

func (app *application) basicAuthUnauthorized(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf(logTemplate, "Basic Auth Unauthorized", r.Method, r.URL.Path, err.Error())
	
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	
	writeJSONError(w, http.StatusUnauthorized, err.Error())
}
