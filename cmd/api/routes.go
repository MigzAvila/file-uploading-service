// Filename cmd/api/routes

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.MethodNotAllowedReponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/randomized/:id", app.randomizedHandler)
	router.HandlerFunc(http.MethodGet, "/v1/personal-info/:id", app.showInfoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/personal-info", app.personalInfo)

	return router
}
