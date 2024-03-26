package main

import (
	"bufio"
	"fmt"
	"hello-go/grettings"
	"os"
	"strings"
)

func main() {
	name, err := askForName()
	if err != nil {
		handleError(err)
	}

	sayHelloTo := grettings.SayHelloWith(grettings.StyleDecorate)
	toDisplay, err := sayHelloTo(name)
	if err != nil {
		handleError(err)
	}

	fmt.Println(toDisplay)
}

func askForName() (string, error) {
	fmt.Println("What is your name ?")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error on asking for name %w", err)
	}

	return strings.TrimSuffix(input, "\n"), nil
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
