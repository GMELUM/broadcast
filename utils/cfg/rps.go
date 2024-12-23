package cfg

import (
	"bufio"
	"errors"
	"fmt"
	"broadcast/config"
	"os"
	"strconv"
	"strings"
)

// rps устанавливает лимит RPS (запросов в секунду).
func rps() error {

	if config.RPSLimit == 0 {

		fmt.Println("Please enter the rps limit (default: 20):")

		reader := bufio.NewReader(os.Stdin)
		rpsValue, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read rps: %v", err.Error())
		}

		// Удаляем пробельные символы (включая \n)
		rpsValue = strings.TrimSpace(rpsValue)

		if rpsValue == "" {
			config.RPSLimit = 20
			return nil
		}

		// Преобразуем строку в число
		value, err := strconv.Atoi(rpsValue)
		if err != nil || value < 1 || value > 30 {
			return errors.New("incorrect rps value. Enter a number between 1 and 30")
		}

		config.RPSLimit = value
	}

	return nil
}
