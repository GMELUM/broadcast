package utils

import (
	"bufio"
	"fmt"
	"os"
)

func HandlerExit() {
	if r := recover(); r != nil {
		fmt.Printf("Program crashed with panic: %v\n", r)
	}

	fmt.Println("Press any key to exit...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	
}
