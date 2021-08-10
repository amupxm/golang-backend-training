package arg

import (
	"flag"
)

var (
	Fi  string
	Shh bool
)

func init() {
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
