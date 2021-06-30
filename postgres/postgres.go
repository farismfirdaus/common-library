package postgres

import (
	"database/sql"

	"github.com/pkg/errors"
)

// DB is PostgreSQL's client
var DB *sql.DB

// Initialize Postgre
func Init(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(errors.Wrap(err, "open postgres: "))
	}

	if err = DB.Ping(); err != nil {
		panic(errors.Wrap(err, "ping postgres: "))
	}
}
