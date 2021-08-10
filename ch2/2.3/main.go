package main

import "fmt"

func main() {
	InitializeArgs()
	if !Shh {
		if Fi == "" {
			fmt.Println("please use --name=YourNameHere")
			return
		}
		fmt.Printf("Hello %s !", Fi)
	}
}
