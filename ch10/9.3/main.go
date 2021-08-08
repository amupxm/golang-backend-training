package main

import (
	"fmt"

	"github.com/amupxm/golang-backend-training/ch10/9.3/srv/money"
)

func main() {
	c, _ := money.ParseCAD("$100.01")
	fmt.Println(c)
	fmt.Printf("money = %#v\n", c)
}
