package grettings

import "fmt"

type StyleDecoration struct{}

func (d StyleDecoration) ApplyTo(value string) string {
	return fmt.Sprintf("\x1B[1m%s, decorated with style\x1B[0m", value)
}
