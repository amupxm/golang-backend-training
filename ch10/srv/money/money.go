package money

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/amupxm/golang-backend-training/ch10/config"
)

type (
	CAD struct {
		Dollar int64
		Cents  int64
	}
)

func Cents(i int64) CAD {
	c := CAD{
		Dollar: i / 100,
		Cents:  i % 100,
	}

	return c
}

func ParseCAD(s string) (*CAD, error) {

	acc := float64(1)
	isNeg := false
	if strings.Contains(s, "-") {
		isNeg = true
	}
	if (!strings.Contains(s, "¢") || !strings.Contains(s, "Cents")) && strings.Contains(s, ".") {
		acc = 100
	}
	// remove all space
	s = strings.Replace(s, " ", "", -1)
	// remove all ,
	s = strings.Replace(s, ",", "", -1)

	// remove all non-numeric characters
	re := regexp.MustCompile(`([-]|\.)?\d[\d,]*[\.]?[\d{2}]*`)
	submatchall := re.FindAllString(s, -1)
	if len(submatchall) != 1 {
		return nil, errors.New(config.InvalidMoneyString)
	}
	s = submatchall[0]
	// if string starts with . add zero to the front
	if strings.HasPrefix(s, ".") {
		s = "0" + s
	}
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, errors.New(config.InvalidMoneyString)
	}
	c := int64(t * acc)
	res := Cents(c)
	if isNeg && (res.Dollar > 0 || res.Cents > 0) {
		res.Dollar = -res.Dollar
		res.Cents = -res.Cents
	}
	return &res, nil
}

func (c *CAD) Abs(o CAD) CAD {
	if c.Dollar < 0 || c.Cents < 0 {
		return o
	}
	return *c
}

func (c *CAD) Add(o CAD) CAD {
	sum := Cents(c.AsCent() + o.AsCent())
	return sum
}

func (c *CAD) AsCent() int64 {
	return c.Cents + c.Dollar*100
}

func (c *CAD) CanonicalForm() (dollar int64, cent int64) {
	return c.Dollar, c.Cents
}

func (c *CAD) Mul(scalar int64) CAD {
	return Cents(c.AsCent() * scalar)
}

func (c *CAD) Sub(o CAD) CAD {
	return Cents(c.AsCent() - o.AsCent())
}
