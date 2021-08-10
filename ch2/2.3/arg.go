package main

import (
	"flag"
)

var (
	Fi  string
	Shh bool
)

func InitializeArgs() {
	flag.StringVar(
		&Fi,
		"name",
		"",
		"name of user to greeting",
	)
	flag.BoolVar(
		&Shh,
		"shhh",
		false,
		"user --shhh flag to mute stdout",
	)
	flag.Parse()

}
