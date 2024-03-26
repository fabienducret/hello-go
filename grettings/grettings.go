package grettings

import "fmt"

type Decorate func(value string) string

type GetNameRepository func() (string, error)

func SayHelloWith(decorate Decorate, getName GetNameRepository) func() (string, error) {
	sayHelloTo := func(name string) string {
		hello := fmt.Sprintf("Hello %s", name)

		return decorate(hello)
	}

	return func() (string, error) {
		name, err := getName()
		if err != nil {
			return "", fmt.Errorf("SayHello: error on getName")
		}

		if name == "" {
			return "", fmt.Errorf("SayHello: empty name")
		}

		return sayHelloTo(name), nil
	}
}
