package grettings

import "fmt"

func StyleDecorate(value string) string {
	return fmt.Sprintf("\x1B[1m%s, decorated with style\x1B[0m", value)
}
