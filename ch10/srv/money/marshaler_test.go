package money

import (
	"fmt"
	"testing"
)

func (c *CAD) TestMarshalJSON(t *testing.T) {
	testTable := []struct {
		M *CAD
		S string
	}{
		{M: &CAD{1, 0}, S: `"CAD$1.00"`},
		{M: &CAD{-11, 0}, S: `"CAD$-11.00"`},
		{M: &CAD{1, 01}, S: `"CAD$1.01"`},
		{M: &CAD{1, 1}, S: `"CAD$1.10"`},
	}
	for testIndex, test := range testTable {
		b, err := test.M.MarshalJSON()
		if err != nil {
			t.Errorf("test %d: with error %v", testIndex, err)
		}
		if string(b) != fmt.Sprintf("\"%s\"", test.S) {
			t.Errorf("test %d: expected %v, got %v", testIndex, test.S, string(b))
		}
	}

}

func (c *CAD) TestUnmarshalJSON(t *testing.T) {

	testTable := []struct {
		M     *CAD
		Bytes []byte
	}{
		{M: &CAD{1, 0}, Bytes: []byte(`"CAD$1.00"`)},
		{M: &CAD{-11, 0}, Bytes: []byte(`"CAD$-11.00"`)},
		{M: &CAD{1, 01}, Bytes: []byte(`"CAD$1.01"`)},
		{M: &CAD{1, 1}, Bytes: []byte(`"CAD$1.10"`)},
	}
	for testIndex, test := range testTable {
		var m *CAD
		err := m.UnmarshalJSON(test.Bytes)
		if err != nil {
			t.Errorf("test %d: with error %v", testIndex, err)
		}
		if m.Cents != test.M.Cents || m.Dollar != test.M.Dollar {
			t.Errorf("test %d: expected %v, got %v", testIndex, test.M, m)
		}
	}
}
