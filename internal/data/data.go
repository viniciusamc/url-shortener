package data

import (
	"database/sql"
)

type Models struct {
	Url UrlModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		Url: UrlModel{db},
	}
}
