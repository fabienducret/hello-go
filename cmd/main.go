package main

import (
	"fmt"
	"hello-go/grettings"
	"os"
)

func main() {
	sayHello := grettings.SayHelloWith(grettings.StyleDecorate, grettings.GetNameInConsole)

	toDisplay, err := sayHello()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(toDisplay)
}
