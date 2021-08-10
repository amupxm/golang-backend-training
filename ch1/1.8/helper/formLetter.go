package helper

import (
	"fmt"
)

type (
	FormLetter interface {
		Print(Letter, Name, Weather, snackName string)
	}
	formLetter struct{}
)

func NewFormLetter() FormLetter {
	return &formLetter{}
}

func (f *formLetter) Print(Letter, Name, Weather, snackName string) {
	fmt.Printf(Letter, Name, Weather, snackName)
}
