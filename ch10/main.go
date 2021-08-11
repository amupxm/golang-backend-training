package main

import (
	"encoding/json"
	"fmt"

	"github.com/amupxm/golang-backend-training/ch10/srv/money"
)

type msg struct {
	Name    string    `json:"name"`
	Balance money.CAD `json:"balance"`
}

func main() {
	c, e := money.ParseCAD("-5")
	if e != nil {
		panic(e)
	}
	fmt.Println(c)
	fmt.Printf("money = %#v\n", c)
	result := &msg{
		Name:    "John Doe",
		Balance: money.Cents(12345),
	}

	o, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(o))
	result = &msg{}
	json.Unmarshal(o, result)
	fmt.Println(result.Balance)
}
