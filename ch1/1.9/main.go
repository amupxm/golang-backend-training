package main

import (
	"fmt"

	config "github.com/amupxm/golang-backend-training/c1/1.8/cfg"
	helper "github.com/amupxm/golang-backend-training/c1/1.8/lib/formLetter"
)

func main() {

	helper.PrintIt(
		fmt.Sprintf(
			config.TheLetter,
			config.Name,
			config.Weather,
			config.Snack,
		),
	)
}
