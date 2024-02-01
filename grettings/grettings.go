package grettings

import "fmt"

type Decoration interface {
	ApplyTo(value string) string
}

func SayHelloWith(decoration Decoration) func(string) (string, error) {
	sayHelloTo := func(name string) string {
		hello := fmt.Sprintf("Hello %s", name)

		return decoration.ApplyTo(hello)
	}

	return func(name string) (string, error) {
		if name == "" {
			return "", fmt.Errorf("empty name")
		}

		return sayHelloTo(name), nil
	}
}
