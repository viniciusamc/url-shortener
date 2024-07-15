package main

import (
	"fmt"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message, "status": status}

	err := app.writeJson(w, status, env, nil)
	if err != nil {
		app.writeJson(w, 500, nil, nil)
	}
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error){
	app.errorResponse(w, r, http.StatusNotFound, err.Error())
}

func (app *application) internalError(w http.ResponseWriter, r *http.Request, err error){
	fmt.Print(err.Error())

	app.errorResponse(w, r, http.StatusInternalServerError, "Try again later")
}
