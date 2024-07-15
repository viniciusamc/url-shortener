package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"url-shortener/internal/data"
)

func (app *application) newURL(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Link     string `json:"link"`
		Duration string `json:"duration"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	timeConv, err := time.Parse("2006-01-02 15:04:05", input.Duration)

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	link, err := app.models.Url.Insert(input.Link, timeConv)
	if err != nil {
		switch {
		case errors.Is(err, data.DuplicateNewLink):
			app.internalError(w, r, err)
			return
		default:
			app.internalError(w, r, err)
			return
		}
	}

	fullURL := fmt.Sprintf("%s://%s/%s", "http", r.Host, link)

	app.writeJson(w, 201, envelope{"shortened_url": fullURL}, nil)
}

func (app *application) goToUrl(w http.ResponseWriter, r *http.Request) {
	words := r.PathValue("words")
	plus := strings.Contains(words, "+")

	if plus {
		words = words[:len(words)-1]
	}

	link, err := app.models.Url.Get(words)
	if err != nil {
		switch {
		case errors.Is(err, data.NotFound):
			app.notFound(w, r, err)
			return
		case errors.Is(err, data.Expired):
			app.notFound(w, r, err)
			return
		default:
			app.internalError(w, r, err)
			return
		}
	}

	headers := http.Header{}

	if plus {
		app.writeJson(w, 200, envelope{"link": link.Link, "url": link.NewLink}, nil)
		return
	}

	headers.Set("Location", link.Link)
	app.writeJson(w, 307, envelope{"link": link.Link, "url": link.NewLink}, headers)
}
