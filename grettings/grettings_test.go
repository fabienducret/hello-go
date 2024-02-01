package grettings_test

import (
	"fmt"
	"hello-go/grettings"
	"testing"
)

type StubDecoration struct{}

func (d StubDecoration) ApplyTo(value string) string {
	return fmt.Sprintf("--> %s <--", value)
}

func TestGrettings(t *testing.T) {

	t.Run("say hello to name", func(t *testing.T) {
		sayHelloTo := grettings.SayHelloWith(StubDecoration{})
		want := "--> Hello Fabien <--"

		got, _ := sayHelloTo("Fabien")

		if got != want {
			t.Errorf("error on say hello to name, want %s, got %s", want, got)
		}
	})

	t.Run("get error for empty name", func(t *testing.T) {
		sayHelloTo := grettings.SayHelloWith(StubDecoration{})

		_, err := sayHelloTo("")

		if err == nil {
			t.Error("error on get error for empty name")
		}
	})
}
