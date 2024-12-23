package storage

import (
	"fmt"
	"broadcast/config"
	"broadcast/utils"
	"broadcast/utils/context"
)

func FillOutTable(tableName string, dumps [][]string, chunkSize int) error {

	db, err := context.GetSQLite()
	if err != nil {
		return context.ErrSQLLiteNotInitialized
	}

	list := utils.SplitStringArray(tableName, dumps, chunkSize)

	bar := utils.NewProgress(fmt.Sprintf("[%-9v]         ", "import data"), len(list) * config.SplitChunkCount)

	for _, query := range list {
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
		bar.Add(config.SplitChunkCount)
	}

	println("")

	return nil

}
