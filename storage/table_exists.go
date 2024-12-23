package storage

import "broadcast/utils/context"

func TableExists(table string) (bool, error) {

	db, err := context.GetSQLite()
	if err != nil {
		return false, context.ErrSQLLiteNotInitialized
	}

	var count int
	query := `SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?;`
	err = db.QueryRow(query, table).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil

}
