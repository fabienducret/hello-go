package grettings

import "fmt"

type SayHello func(string) (string, error)

type Decorate func(value string) string

func SayHelloWith(decorate Decorate) SayHello {
	sayHelloTo := func(name string) string {
		hello := fmt.Sprintf("Hello %s", name)

		return decorate(hello)
	}

	return func(name string) (string, error) {
		if name == "" {
			return "", fmt.Errorf("empty name")
		}

		return sayHelloTo(name), nil
	}
}
