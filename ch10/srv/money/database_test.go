package money_test

import (
	"database/sql/driver"
	"testing"

	"github.com/amupxm/golang-backend-training/ch10/srv/money"
)

func TestValue(t *testing.T) {
	testTable := []struct {
		S        money.CAD
		Expected driver.Value
	}{
		{money.CAD{Dollar: 1, Cents: 0}, "1.0"},
	}

	for _, test := range testTable {
		actual, err := test.S.Value()
		if actual != test.Expected {
			t.Errorf("Expected %v, got %v", test.Expected, actual)
		}
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	}
}
