// Filename: cmd/api/helpers

package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Convert our map into a JSON object
	// js, err := json.Marshal(data)
	// Format the JSON object for cmd -- Takes more resources than printing it normally
	js, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	// Add a newline to make viewing on the terminal easier
	js = append(js, '\n')

	// Add the headers
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Specify that we will serve our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// Write the I I byte slice containing the JSON response body
	w.Write(js)
	return nil
}

func (app *application) generateRandomString(length int) string {
	s := make([]rune, length)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}

// Utility function for reading ID in Endpoint
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Use the param
	// Use the ParamsFormContext
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}

	return id, nil

}
