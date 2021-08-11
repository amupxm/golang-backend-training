package money

import "encoding/json"

func (c CAD) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *CAD) UnmarshalJSON(b []byte) error {

	parsedCad, err := ParseCAD(string(b))
	if err != nil {
		return err
	}
	*c = *parsedCad
	return nil
}
