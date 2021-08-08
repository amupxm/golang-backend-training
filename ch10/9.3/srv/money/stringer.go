package money

import "fmt"

func (c *CAD) GoString() string {
	return fmt.Sprintf("Cents(%d)", c.AsCent())
}

func (c *CAD) String() string {
	z := int64(1)
	if c.Dollar < 0 || c.Cents < 0 {
		z = -1
	}
	r := fmt.Sprintf("%d", c.Dollar)
	if c.Cents != 0 {
		r += fmt.Sprintf(".%02d", c.Cents*z)
	}
	if z == -1 {
		r = "-" + r
	}
	return "CAD$" + r
}
