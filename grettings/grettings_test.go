package grettings_test

import (
	"fmt"
	"hello-go/grettings"
	"testing"
)

func TestGrettings(t *testing.T) {
	t.Run("say hello to name", func(t *testing.T) {
		getName := func() (string, error) {
			return "Fabien", nil
		}
		sayHello := grettings.SayHelloWith(decorate, getName)
		want := "--> Hello Fabien <--"

		got, _ := sayHello()

		if got != want {
			t.Errorf("error on say hello to name, want %s, got %s", want, got)
		}
	})

	t.Run("get error for empty name", func(t *testing.T) {
		getName := func() (string, error) {
			return "", nil
		}
		sayHello := grettings.SayHelloWith(decorate, getName)

		_, err := sayHello()

		if err == nil {
			t.Error("error on get error for empty name")
		}
	})
}

func decorate(value string) string {
	return fmt.Sprintf("--> %s <--", value)
}
