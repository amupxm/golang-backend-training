package money_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/amupxm/golang-backend-training/ch10/9.3/srv/money"
)

func TestStringer(t *testing.T) {
	testTable := []struct {
		S        string
		Expected string
	}{
		{"1.23", "CAD$1.23"},
		{"-CAD$.2", "CAD$-0.20"},
	}

	for _, tt := range testTable {
		var b bytes.Buffer
		money, _ := money.ParseCAD(tt.S)
		fmt.Fprint(
			&b,
			money,
		)
		if b.String() != tt.Expected {
			t.Error("invalid got : " + b.String() + " expected " + tt.Expected)
		}
	}
}

func TestGoStringer(t *testing.T) {
	testTable := []struct {
		S        string
		Expected string
	}{
		{"1.23", "Cents(123)"},
		{"-CAD$.2", "Cents(-20)"},
	}

	for _, tt := range testTable {
		m, _ := money.ParseCAD(tt.S)
		var b bytes.Buffer
		fmt.Fprintf(
			&b,
			"%#v",
			m,
		)
		if b.String() != tt.Expected {
			t.Error("invalid got : " + b.String() + " expected " + tt.Expected)

		}
	}

}
