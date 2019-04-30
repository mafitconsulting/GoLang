package main

import (
	"GoLang/dog"
	"fmt"
)

func main() {

	var i int
	fmt.Print("Enter Dog Year?\n")
	fmt.Scan(&i)
	fmt.Println("The age of the dog is", dog.Years(i))
}
