package main

import (
	"flag"
	"fmt"
)

func main() {
	var fi string
	var shh bool
	flag.StringVar(
		&fi,
		"name",
		"",
		"name of user to greeting",
	)
	flag.BoolVar(
		&shh,
		"shhh",
		false,
		"user --shhh flag to mute stdout",
	)
	flag.Parse()
	if !shh {
		if fi == "" {
			fmt.Println("please use --name=YourNameHere")
			return
		}
		fmt.Printf("Hello %s !", fi)

	}
}
