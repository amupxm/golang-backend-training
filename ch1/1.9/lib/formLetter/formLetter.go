package helper

import (
	"fmt"
	"strings"
)

func PrintIt(args ...string) {
	fmt.Println(
		strings.Join(args, "\n"),
	)
}
