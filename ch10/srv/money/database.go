package money

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

func (c *CAD) Value() (driver.Value, error) {
	return c.AsCent, nil
}

func (c *CAD) Scan(src interface{}) error {
	var str string
	switch src := src.(type) {
	case []byte:
		str = string(src[0])
	case int64:
		str = fmt.Sprint(src)
	default:
		return errors.New("database data is invalid")
	}
	money, err := ParseCAD(str)
	if err != nil {
		return err
	}
	*c = *money
	return nil
}
