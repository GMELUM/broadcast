package utils

import (
	"fmt"
	"strings"
)

func SplitStringArray(tableName string, dumps [][]string, chunkSize int) []string {

	var queries []string
	length := len(dumps[1:])

	for i := 0; i < length; i += chunkSize {
		end := i + chunkSize
		if end > length {
			end = length
		}

		chunk := dumps[1:][i:end]
		value := strings.Builder{}
		value.WriteString(fmt.Sprintf("INSERT INTO `%v`(`user`) VALUES ", tableName))

		for j, row := range chunk {
			values := "'" + strings.Join(row, "','") + "'"
			if j == len(chunk)-1 {
				value.WriteString(fmt.Sprintf("(%v);", values))
			} else {
				value.WriteString(fmt.Sprintf("(%v),", values))
			}
		}

		queries = append(queries, value.String())
	}

	return queries
}
