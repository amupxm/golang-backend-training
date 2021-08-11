package money

import (
	"errors"
	"fmt"
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
	neg := false
	if i < 0 {
		neg = true
		i = -i
	}
	c := CAD{
		Dollar: i / 100,
		Cents:  i % 100,
	}
	if neg {
		c.Dollar = -c.Dollar
		c.Cents = -c.Cents
	}
	return c
}

func ParseCAD(s string) (*CAD, error) {
	// helper function to parse a string into a int64
	convToInt := func(str string) int64 {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
	// CheckIs negative?
	isNeg := int64(1)
	if strings.Contains(s, "-") {
		isNeg = -1
		s = strings.Replace(s, "-", "", 1)
	}
	// CheckIs all Cents?
	isCents := strings.Contains(s, "¢") || strings.Contains(s, "Cents")
	// remove spaces and commas
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", "", -1)
	// Extract only the numbers
	re := regexp.MustCompile(`([-]|\.)?\d[\d,]*[\.]?[\d{2}]*`)
	submatchall := re.FindAllString(s, -1)
	if len(submatchall) != 1 {
		return nil, errors.New(config.InvalidMoneyString)
	}
	s = submatchall[0]
	// if string starts with . add zero to the front
	if strings.HasPrefix(s, ".") {
		s = fmt.Sprintf("0%s", s)
	}
	// split string by dot
	sArr := strings.Split(s, ".")
	// if string has more than 2 parts, return error
	if len(sArr) > 2 {
		return nil, errors.New(config.InvalidMoneyString)
	}
	//if its only cents , convert ans return only cents

	if isCents && len(sArr) == 1 {
		fmt.Println("cents", s)
		c := convToInt(sArr[0])
		asCent := Cents(c * isNeg)
		return &asCent, nil
	}

	//if itd dollas , convert to cents and return

	if !isCents && len(sArr) == 2 {
		d := convToInt(sArr[0])
		// limit c len to 2
		scalar := 0
		// count of 0 in begining of sArr[1]
		for _, c := range sArr[1] {
			if c == '0' {
				scalar++
			} else {
				break
			}
		}
		if scalar > 2 {
			sArr[1] = "0"
		}
		if scalar == 0 {
			sArr[1] = fmt.Sprintf("%s0", sArr[1])

		}
		if len(sArr[1]) > 2 {
			sArr[1] = sArr[1][:2]
		}
		c := convToInt(sArr[1])
		dollars := CAD{Dollar: d * isNeg, Cents: (c * isNeg)}

		return &dollars, nil
	}
	if !isCents && len(sArr) == 1 {
		dollar := CAD{Dollar: convToInt(sArr[0]) * isNeg, Cents: 0}
		return &dollar, nil
	}

	return nil, errors.New(fmt.Sprint(isCents, len(sArr), s))
}

// func ParseCAD(s string) (*CAD, error) {

// 	acc := int64(1)
// 	isNeg := false
// 	if strings.Contains(s, "-") {
// 		isNeg = true
// 	}
// 	if (!strings.Contains(s, "¢") || !strings.Contains(s, "Cents")) && strings.Contains(s, ".") {
// 		acc = 100
// 	}
// 	// remove all space
// 	s = strings.Replace(s, " ", "", -1)
// 	// remove all ,
// 	s = strings.Replace(s, ",", "", -1)

// 	// remove all non-numeric characters
// 	re := regexp.MustCompile(`([-]|\.)?\d[\d,]*[\.]?[\d{2}]*`)
// 	submatchall := re.FindAllString(s, -1)
// 	if len(submatchall) != 1 {
// 		return nil, errors.New(config.InvalidMoneyString)
// 	}
// 	s = submatchall[0]
// 	// if string starts with . add zero to the front
// 	if strings.HasPrefix(s, ".") {
// 		s = "0" + s
// 	}
// 	sep := strings.Split(s, ".")
// 	if len(sep) > 2 || len(sep) == 0 {
// 		return nil, errors.New(config.InvalidMoneyString)
// 	}
// 	if len(sep) == 1 {
// 		sep = []string{"0", sep[0]}
// 	}
// 	convToInt := func(str string) int64 {
// 		i, err := strconv.ParseInt(str, 10, 64)
// 		if err != nil {
// 			return 0
// 		}
// 		return i
// 	}
// 	// set max len sep[1] = 2
// 	if len(sep[1]) > 2 {
// 		sep[1] = sep[1][:2]
// 	}

// 	c := convToInt(sep[0])*acc + (convToInt(sep[1]))
// 	// c := int64(t * acc)
// 	res := Cents(c)
// 	if isNeg && (res.Dollar > 0 || res.Cents > 0) {
// 		res.Dollar = -res.Dollar
// 		res.Cents = -res.Cents
// 	}
// 	return &res, nil
// }

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
