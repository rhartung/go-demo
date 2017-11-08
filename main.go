package main

import (
	"fmt"
)

func main() {

	// Variable declaration and assignment - shorthand
	// Variables accessible within function
	// can only use shorthand notation inside func
	a := "Oh hey there"
	b := 4
	c := true
	d := 8.91

	// %v used for string formatting w/variables
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	// %T used for formatting string to show variable type
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
