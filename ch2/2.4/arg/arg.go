package arg

import (
	"flag"
	"fmt"
)

func InitArgs() {
	var fi string
	var shh bool
	flag.StringVar(
		&fi,
		"name",
		"EMPTY_NAME",
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
		if fi == "EMPTY_NAME" {
			fmt.Println("please use --name flag like :\n go run main.go --name=YourNameHere")
			return
		}
		fmt.Println("Hello " + fi + "!")
	}
}
