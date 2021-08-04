package main

import "fmt"

type (
	FunIn struct {
		Name    string
		Weather string
		Snack   string
	}
)

func PrintIt(i *FunIn) {
	fmt.Printf("Hello %s!\nThe weather today is %s.\nToday's snack will be %s.", i.Name, i.Weather, i.Snack)
}
func main() {
	f := FunIn{"Bob", "cloudy", "chips"}
	PrintIt(&f)
}
