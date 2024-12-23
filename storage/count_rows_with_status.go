package storage

import (
	"broadcast/utils/context"
)

// CountRowsWithStatus counts the number of rows in the queue table where status.
func CountRowsWithStatus(status int) (int, error) {
	// Get the database connection from the context
	db, err := context.GetSQLite()
	if err != nil {
		return 0, context.ErrSQLLiteNotInitialized
	}

	// // Prepare the SQL query to count rows with status = 0
	// query := fmt.Sprintf(`SELECT COUNT(*) FROM queue WHERE status = %v;`, status)

    // Prepare the SQL query to count rows with status = 0
    query := `SELECT COUNT(*) FROM queue`

	// Variable to store the count result
	var count int

	// Execute the query and scan the result into the count variable
	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err // Return the error if the query fails
	}

	return count, nil
}
