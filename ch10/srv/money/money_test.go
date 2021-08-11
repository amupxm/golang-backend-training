package money_test

import (
	"testing"

	"github.com/amupxm/golang-backend-training/ch10/srv/money"
)

func TestCents(t *testing.T) {
	testTable := []struct {
		I      int64
		Result money.CAD
	}{
		{0, money.CAD{0, 0}},
		{-0, money.CAD{0, 0}},
		{105, money.CAD{1, 5}},
		{1050, money.CAD{10, 50}},
		{-105, money.CAD{-1, -5}},
		{-1, money.CAD{0, -1}},
		{-500, money.CAD{-5, 0}},

		{-100000005, money.CAD{-1000000, -05}},
		{-100000005, money.CAD{-1000000, -05}},
	}
	for _, test := range testTable {
		if got := money.Cents(test.I); got != test.Result {
			t.Errorf("Cents(%d) = %v; want %v", test.I, got, test.Result)
			continue
		}
	}
}

func TestParseCAD(t *testing.T) {
	testTable := []struct {
		S      string
		Result money.CAD
	}{
		{
			"CAD-$1,234.56",
			money.CAD{
				-1234,
				-56,
			},
		},
		{
			"$-.09",
			money.CAD{
				0,
				-9,
			},
		},
		{
			"CAD$-.9",
			money.CAD{
				0,
				-90,
			},
		}, {
			"0.02",
			money.CAD{
				0,
				2,
			},
		},
		{
			"-5",
			money.CAD{
				-5,
				0,
			},
		},
		{
			"CAD$-.09",
			money.CAD{
				0,
				-9,
			},
		},
		{
			"9¢",
			money.CAD{
				0,
				9,
			},
		},
		{
			"-9¢",
			money.CAD{
				0,
				-9,
			},
		},
		{
			"-$1235.56",
			money.CAD{
				-1235,
				-56,
			},
		},
		{
			"-123456¢",
			money.CAD{
				-1234,
				-56,
			},
		},
		{
			"123456¢",
			money.CAD{
				1234,
				56,
			},
		},
		{
			"$.000123456",
			money.CAD{
				0,
				0,
			},
		},
	}
	for _, test := range testTable {
		got, err := money.ParseCAD(test.S)
		if err != nil {
			t.Error(err)
		}
		if *got != test.Result {
			t.Errorf(test.S, got.AsCent(), test.Result.AsCent())
			continue
		}
	}
}

func TestAdd(t *testing.T) {
	testTable := []struct {
		A      money.CAD
		B      money.CAD
		Result money.CAD
	}{
		{
			money.CAD{1, 0},
			money.CAD{1, 0},
			money.CAD{2, 0},
		},
		{
			money.CAD{0, 0},
			money.CAD{0, 0},
			money.CAD{0, 0},
		},
		{
			money.CAD{0, 00},
			money.CAD{0, 2},
			money.CAD{0, 2},
		},
		{
			money.CAD{0, 11},
			money.CAD{0, 22},
			money.CAD{0, 33},
		},
	}
	for _, test := range testTable {
		if got := test.A.Add(test.B); got != test.Result {
			t.Errorf("Add(%v, %v) = %v; want %v", test.A, test.B, got, test.Result)
			continue
		}
	}
}

func TestMultiply(t *testing.T) {
	testTable := []struct {
		A      money.CAD
		B      int64
		Result money.CAD
	}{
		{
			money.CAD{1, 0},
			1,
			money.CAD{1, 0},
		},
		{
			money.CAD{0, 0},
			0,
			money.CAD{0, 0},
		},
		{
			money.CAD{1, 1},
			2,
			money.CAD{2, 2},
		},
	}
	for _, test := range testTable {
		if got := test.A.Mul(test.B); got != test.Result {
			t.Errorf("Sub(%v, %v) = %v; want %v", test.A, test.B, got, test.Result)
			continue
		}
	}
}

func TestSub(t *testing.T) {
	testTable := []struct {
		A      money.CAD
		B      money.CAD
		Result money.CAD
	}{
		{
			money.CAD{1, 0},
			money.CAD{1, 0},
			money.CAD{0, 0},
		},
		{
			money.CAD{0, 0},
			money.CAD{1, 0},
			money.CAD{-1, 0},
		},
		{
			money.CAD{4, 1},
			money.CAD{1, 0},
			money.CAD{3, 1},
		},
	}
	for _, test := range testTable {
		if got := test.A.Sub(test.B); got != test.Result {
			t.Errorf("Sub(%v, %v) = %v; want %v", test.A, test.B, got, test.Result)
			continue
		}
	}
}
