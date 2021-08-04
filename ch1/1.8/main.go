package main

import (
	"fmt"

	"github.com/amupxm/golang-backend-training/c1/1.8/config"
	"github.com/amupxm/golang-backend-training/c1/1.8/helper"
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
