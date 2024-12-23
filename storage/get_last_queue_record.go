package storage

import (
	"database/sql"
	"fmt"
	"broadcast/utils/context"
)

// Queue represents a record in the queue table.
type Queue struct {
	ID     int  `json:"id"`
	User   int  `json:"user"`
	Status int  `json:"success"`
}

// GetLastQueueRecord retrieves the last record where status > 0 as a Queue struct.
func GetLastQueueRecord() (*int, error) {
	// Get the database connection from the context
	db, err := context.GetSQLite()
	if err != nil {
		return nil, context.ErrSQLLiteNotInitialized
	}

	// Prepare the query to get the last record with status > 0
	query := `SELECT COUNT(*) FROM queue WHERE status = 0;`

	// Create a variable to store the result
	var record int

	// Execute the query and scan the result into the struct
	err = db.QueryRow(query).Scan(&record)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no rows are found, return a specific message
			return nil, fmt.Errorf("no record found with status > 0")
		}
		return nil, fmt.Errorf("failed to query the last record: %w", err)
	}

	return &record, nil
}
