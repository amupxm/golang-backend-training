package main

import (
	"fmt"

	"github.com/amupxm/golang-backend-training/c2/2.2/arg"
)

func main() {
	if !arg.Shh {
		if arg.Fi == "" {
			fmt.Println("please use --name=YourNameHere")
			return
		}
		fmt.Printf("Hello %s !", arg.Fi)
	}
}
