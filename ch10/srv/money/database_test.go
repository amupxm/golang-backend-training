package money_test

import (
	"testing"

	"github.com/amupxm/golang-backend-training/ch10/srv/money"
)

func TestScan(t *testing.T) {
	testTable := []struct {
		Value    interface{}
		Expected money.CAD
	}{
		{
			Value:    -5,
			Expected: money.Cents(-500),
		},
		{
			Value:    -4,
			Expected: money.Cents(-400),
		},
		{
			Value:    -3,
			Expected: money.Cents(-300),
		},
		{
			Value:    -2,
			Expected: money.Cents(-200),
		},
		{
			Value:    -1,
			Expected: money.Cents(-100),
		},
		{
			Value:    0,
			Expected: money.Cents(0),
		},
		{
			Value:    1,
			Expected: money.Cents(100),
		},
		{
			Value:    2,
			Expected: money.Cents(200),
		},
		{
			Value:    byte(3),
			Expected: money.Cents(300),
		},
		{
			Value:    "4",
			Expected: money.Cents(400),
		},
		{
			Value:    5,
			Expected: money.Cents(500),
		},

		{
			Value:    "0.01",
			Expected: money.Cents(1),
		},
		{
			Value:    "0.02",
			Expected: money.Cents(2),
		},
	}

	for testNumber, test := range testTable {
		var m money.CAD
		clone := test.Value
		err := m.Scan(&clone)

		if err != nil {
			t.Errorf("number %d :: Error on converting to money %v with error %s", testNumber+1, test.Value, err.Error())
			continue
		} else {
			if m.Dollar != test.Expected.Dollar || m.Cents != test.Expected.Cents {
				t.Errorf("number %d :: Expected %v got %v", testNumber+1, test.Expected, m)
				continue
			}
		}

	}
}

func TestValue(t *testing.T) {
	testTable := []struct {
		Value    money.CAD
		Expected string
	}{
		{
			Value:    money.Cents(10001),
			Expected: "100.1",
		},
		{
			Value:    money.Cents(-500),
			Expected: "-5.0",
		},
		{
			Value:    money.Cents(-400),
			Expected: "-4.0",
		},
		{
			Value:    money.Cents(-300),
			Expected: "-3.0",
		},
		{
			Value:    money.Cents(-200),
			Expected: "-2.0",
		},
		{
			Value:    money.Cents(-100),
			Expected: "-1.0",
		},
		{
			Value:    money.Cents(0),
			Expected: "0.0",
		},
		{
			Value:    money.Cents(100),
			Expected: "1.0",
		},
	}
	for testNumber, test := range testTable {
		v, err := test.Value.Value()
		if err != nil {
			t.Errorf("number %d :: Error on converting to string %v with error %s", testNumber+1, test.Value, err.Error())
			continue
		}
		if v != test.Expected {
			t.Errorf("number %d :: Expected %s got %s", testNumber+1, test.Expected, v)
			continue
		}
	}
}
