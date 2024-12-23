package storage

import (
	"database/sql"
	"broadcast/models"
	"broadcast/utils/context"
	"broadcast/utils/sqlite"
)

// GetUsers возвращает список пользователей, фильтруя по id в указанном диапазоне (page и limit).
func GetUsers(page int, limit int, idStart int64) (*[]models.Queue, error) {

	ctx, err := context.GetSQLite()
	if err != nil {
		return nil, context.ErrSQLLiteNotInitialized
	}

	start := (page * limit) + int(idStart)
	end := start + limit - 1

	// Создаём SQL-запрос с параметрами
	query := "SELECT `id`, `user`, `status` FROM `queue` WHERE id >= ? AND id <= ?"

	// Выполняем запрос
	return sqlite.Query(ctx, sqlite.Params{
		Query: query,
		Args:  []interface{}{start, end},
	}, func(rows *sql.Rows) (*[]models.Queue, error) {
		users := []models.Queue{}

		for rows.Next() {
			user := models.Queue{}
			err := rows.Scan(
				&user.ID,
				&user.User,
				&user.Status,
			)

			if err != nil {
				return nil, err
			}

			users = append(users, user)
		}

		return &users, nil
	})
}
