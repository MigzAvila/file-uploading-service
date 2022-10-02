// Filename: cmd/api/healthcheck.go

package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map to healthcheck requests
	data := envelope{
		"status": "available",
		"system_info": envelope{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
}
