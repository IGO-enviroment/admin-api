package students

import (
	"admin-api/config"
	"database/sql"
)

type Service struct {
	pg       *sql.DB
	settings config.Settings
}

func NewStudentsService(pg *sql.DB, settings config.Settings) Service {
	return Service{
		pg:       pg,
		settings: settings,
	}
}
