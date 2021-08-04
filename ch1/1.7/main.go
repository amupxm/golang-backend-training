package main

import (
	"fmt"

	"github.com/amupxm/golang-backend-training/c1/1.7/config"
)

type (
	FunIn struct {
		Name    string
		Weather string
		Snack   string
	}
)

func PrintIt(i *FunIn) {
	fmt.Printf(config.Letter, i.Name, i.Weather, i.Snack)
}
func main() {
	f := FunIn{
		config.Name,
		config.Weather,
		config.Snack,
	}
	PrintIt(&f)
}
