package cfg

import (
	"bufio"
	"errors"
	"fmt"
	"broadcast/config"
	"os"
	"regexp"
	"strings"
)

func token() error {

	if config.Token == "" {

		println("Token not defined. Please enter the token:")

		reader := bufio.NewReader(os.Stdin)
		token, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read token: %v", err.Error())
		}

		value := strings.TrimSpace(token)

		matched, err := regexp.MatchString(`^[0-9]{8,10}:[a-zA-Z0-9-_]{35}$`, value)
		if err != nil || !matched {
			return errors.New("incorrect format token")
		}

		config.Token = value

	}

	return nil

}
