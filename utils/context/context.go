package context

import (
	"errors"
	"broadcast/utils/sqlite"
)

var (
	// ms holds the MySQL core instance for database operations.
	msl *sqlite.Core
)

// Predefined error messages
var (
	ErrSQLLiteNotInitialized = errors.New("sqlite is not initialized")
)

// SetSQLite sets the MySQL core instance.
// It is used to initialize the `my` variable globally.
//
// Parameters:
//
// - core: A pointer to the mysql.Core instance to be set.
func SetSQLite(core *sqlite.Core) {
	msl = core
}

// GetSQLite returns the initialized MySQL core instance.
// If the instance is not set, it returns an error.
func GetSQLite() (*sqlite.Core, error) {
	if msl == nil {
		return nil, ErrSQLLiteNotInitialized
	}
	return msl, nil
}

// Clear resets all global variables to nil.
// This is useful for cleanup in testing or application shutdown.
func Clear() {
	SetSQLite(nil)
}
