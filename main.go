package main

import (
	"fmt"
)

func main() {
	input := "@bob @john (success) such a cool feature;\nhttps://twitter.com/jdorfman/status/430511497475670016"
	fmt.Println("Input: ", input)
	output := Parse(input)
	fmt.Print("Output: ", output)
}
