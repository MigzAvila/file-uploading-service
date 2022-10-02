package main

import "net/http"

func (app *application) randomizedHandler(w http.ResponseWriter, r *http.Request) {
	//Utilize Utility Methods From helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Extract randomized value based on :id parameter
	randomized_value := app.generateRandomString(int(id))
	err = app.writeJSON(w, http.StatusOK, envelope{"randomized_value": randomized_value}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
