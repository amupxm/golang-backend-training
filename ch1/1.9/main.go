package main

import (
	config "github.com/amupxm/golang-backend-training/chapter1/section1.9/cfg"
	helper "github.com/amupxm/golang-backend-training/chapter1/section1.9/lib/formLetter"
)

func main() {
	helper.NewFormLetter().Print(
		config.TheLetter,
		config.Name,
		config.Weather,
		config.Snack,
	)
}
