package main

import "net/http"

func (app *application) routes() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("GET /health", health)

	r.HandleFunc("POST /url", app.newURL)
	r.HandleFunc("GET /{words}", app.goToUrl)

	return r
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
