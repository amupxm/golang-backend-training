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
		"EMPTY_NAME",
		"name of user to greeting",
	)
	flag.Parse()
	if fi == "EMPTY_NAME" {
		fmt.Println("please use --name flag like :\n go run main.go --name=YourNameHere")
		return
	}
	fmt.Println("Hello " + fi + "!")
}
