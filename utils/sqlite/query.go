package sqlite

import (
	"context"
	"time"

	"database/sql"
)

type Params struct {
	Query string        // SQL query string
	Args  []interface{} // Arguments for the SQL query
}

func Query[T any](
	c *Core,
	params Params,
	callback func(rows *sql.Rows) (*T, error),
) (*T, error) {

	// Create a context with a timeout for the query execution
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel() // Cancel the context after the query execution

	// Execute the query with the provided arguments
	rows, err := c.QueryContext(ctx, params.Query, params.Args...)
	if err != nil {
		// Return the SQL error if it is any other error
		return nil, err
	}
	defer rows.Close() // Close the rows after finishing the query

	// Call the callback function to process the rows and extract the result
	clbRes, clbErr := callback(rows)

	// Return the result and any potential MySQL error from the callback
	return clbRes, clbErr

}
