package money

import (
	"fmt"
)

func (c *CAD) MarshalJSON() ([]byte, error) {
	return []byte(
		fmt.Sprintf("\"%s\"", c.String()),
	), nil
}

func (c *CAD) UnmarshalJSON(b []byte) error {

	parsedCad, err := ParseCAD(string(b))
	if err != nil {
		return err
	}
	*c = *parsedCad
	return nil
}
