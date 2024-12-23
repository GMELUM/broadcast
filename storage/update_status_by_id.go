package storage

import (
	"fmt"
	"broadcast/utils/context"
)

// UpdateStatusByID updates the status field for a specific record in the queue table by ID.
func UpdateStatusByID(id int, status int) error {
	// Get the database connection from the context
	db, err := context.GetSQLite()
	if err != nil {
		return context.ErrSQLLiteNotInitialized
	}

	// Prepare the SQL query for updating the status
	query := `UPDATE queue SET status = ? WHERE id = ?;`

	// Execute the query with the provided status and ID
	result, err := db.Exec(query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update status for ID %d: %w", id, err)
	}

	// Check the number of affected rows to ensure the update was successful
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no record found with ID %d", id)
	}

	return nil
}
