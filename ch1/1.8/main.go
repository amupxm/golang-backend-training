package main

import (
	"github.com/amupxm/golang-backend-training/chapter1/section1.8/config"
	"github.com/amupxm/golang-backend-training/chapter1/section1.8/helper"
)

func main() {
	helper.NewFormLetter().Print(
		config.TheLetter,
		config.Name,
		config.Weather,
		config.Snack,
	)
}
