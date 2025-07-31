package main

import (
	"net/http"
	"time"
)

type status struct {
	Status  string `json:"status"`
	Env     string `json:"env"`
	Version string `json:"version"`
}

// HealthCheck godoc
//
//	@Summary		Health check endpoint
//	@Description	Returns the status of the API service
//	@Produce		json
//	@Success		200	{object}	status	"Returns service status information"
//	@Failure		500	{string}	string
//	@Router			/v1/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	
	time.Sleep(time.Second * 4)

	s := status{
		"ok",
		app.config.env,
		version,
	}

	if err := app.jsonResponse(w, http.StatusOK, s); err != nil {
		app.internalServerError(w, r, err)
	}
}
