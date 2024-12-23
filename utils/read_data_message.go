package utils

import (
	"encoding/json"
	"io"
	"broadcast/models"
	"os"
)

func ReadDataMessage(filePath string) (*models.Message, error) {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Создаем переменную для хранения данных структуры
	var message models.Message

	// Парсим JSON в структуру
	err = json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}