package main

import (
	"net/http"
	"time"

	"fileuploading.miguelavila.net/internal/data"
	"fileuploading.miguelavila.net/internal/validator"
)

// Create a New user
func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Target decode destination

	var input struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Email   string `json:"email"`
		Address string `json:"address"`
		School  string `json:"school"`
		Degree  string `json:"degree"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badResquestReponse(w, r, err)
		return
	}
	// Copy the values from the input struct to a new user struct
	user := &data.User{
		ID:      1,
		Name:    input.Name,
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
		School:  input.School,
		Degree:  input.Degree,
		Version: 1,
	}

	// Initialize a new instance of validator
	v := validator.New()

	// Check the errors maps if there were any errors validation
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Display valid input
	err = app.writeJSON(w, http.StatusOK, envelope{"user_info": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// GET User /v1/personalinfo endpoints
func (app *application) showInfoHandler(w http.ResponseWriter, r *http.Request) {
	// sample data
	user := data.User{
		ID:        12,
		CreatedAt: time.Now(),
		Name:      "Miguel Avila",
		Phone:     "601-4411",
		Email:     "amiguelavila01@gmail.com",
		Address:   "16 Victoria Street",
		School:    "UB",
		Degree:    "AINT",
		Version:   1,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"user_info": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
