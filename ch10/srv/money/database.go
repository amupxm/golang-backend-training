package money

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

func (c CAD) Value() (driver.Value, error) {
	return fmt.Sprintf("%d.%d", c.Dollar, c.Cents), nil
}

func (c *CAD) Scan(src interface{}) error {
	var str string
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.String {
			str = v.Interface().(string)
		}
	case reflect.String:
		str = v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str = fmt.Sprintf("%d", v.Uint())
	default:
		return fmt.Errorf("unsupported type: %s", v.Kind())
	}
	m, err := ParseCAD(str)
	if err != nil {
		fmt.Println(str)
		return err
	}
	c.Dollar = m.Dollar
	c.Cents = m.Cents
	return nil
}
