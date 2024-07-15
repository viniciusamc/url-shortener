package main

import (
	"database/sql"
	"net/http"
	"url-shortener/internal/data"

	_ "github.com/mattn/go-sqlite3"
)

type configs struct {
	port string
}

type application struct {
	config configs
	models data.Models
}

func main() {

	var config configs

	config.port = "8123"

	db := openDB()

	app := application{
		config: config,
		models: data.NewModel(db),
	}

	println("Running on port", config.port)
	http.ListenAndServe(":"+config.port, app.routes())
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./internal/database/database.sqlite3")
	if err != nil {
		panic("DB NOT FOUND")
	}

	err = db.Ping()
	if err != nil {
		panic("DB PING NOT SUCCESSFUL")
	}

	return db
}
