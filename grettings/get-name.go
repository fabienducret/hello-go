package grettings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetNameInConsole() (string, error) {
	fmt.Println("What is your name ?")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error on asking for name %w", err)
	}

	return strings.TrimSuffix(input, "\n"), nil
}
