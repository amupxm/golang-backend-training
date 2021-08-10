package main

import (
	"flag"
	"fmt"
)

func main() {
	var fi string
	flag.StringVar(
		&fi,
		"name",
		"",
		"name of user to greeting",
	)
	flag.Parse()
	if fi == "" {
		fmt.Println("please use --name=YourNameHere")
		return
	}
	fmt.Printf("Hello %s !", fi)
}
