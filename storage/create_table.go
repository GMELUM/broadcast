package storage

import (
	"errors"
	"fmt"
	"broadcast/utils/context"
	"strings"
)

func CreateTable() error {
	tableName := "queue"

	// Define table columns
	columns := map[string]string{
		"id":     "INTEGER PRIMARY KEY AUTOINCREMENT",
		"user":   "BIGINT",
		"status": "INT DEFAULT 0",
	}

	// Build the table creation query
	var columnDefs []string
	for columnName, columnType := range columns {
		definition := fmt.Sprintf("`%s` %s", columnName, columnType)
		columnDefs = append(columnDefs, definition)
	}
	createTableQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s);", tableName, strings.Join(columnDefs, ", "))

	// Build queries for indexes
	createUniqueIndexQuery := fmt.Sprintf("CREATE UNIQUE INDEX IF NOT EXISTS queue_id_uindex ON `%s` (id);", tableName)
	createStatusIndexQuery := fmt.Sprintf("CREATE INDEX IF NOT EXISTS queue_status_index ON `%s` (status);", tableName)

	// Get the database connection
	db, err := context.GetSQLite()
	if err != nil {
		return context.ErrSQLLiteNotInitialized
	}

	// Execute the table creation query
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return errors.New("error creating table")
	}

	// Execute the index creation queries
	_, err = db.Exec(createUniqueIndexQuery)
	if err != nil {
		return errors.New("error creating unique index on id")
	}

	_, err = db.Exec(createStatusIndexQuery)
	if err != nil {
		return errors.New("error creating index on status")
	}

	return nil
}
