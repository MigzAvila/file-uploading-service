package main

import (
	"net/http"
	"time"

	"fileuploading.miguelavila.net/internal/data"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	return
}

// createUserHandler for GET /v1/personal-info endpoints
func (app *application) showInfoHandler(w http.ResponseWriter, r *http.Request) {
	//Utilize Utility Methods From helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the user struct containing the ID we extracted from
	// From URL and sample data
	user := data.User{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Belize Street",
		Phone:     "601-4411",
		Email:     "amiguelavila01@gmail.com",
		Address:   "14 Apple Street",
		School:    "UB",
		Degree:    "AINT",
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user_info": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
