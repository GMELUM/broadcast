package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Core struct {
	*sql.DB
}

func New(path string) (*Core, error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure the connection is valid.
	err = db.Ping()
	if err != nil {
		return nil, err // Terminate the program if the connection is invalid.
	}

	core := &Core{db}

	return core, nil

}
