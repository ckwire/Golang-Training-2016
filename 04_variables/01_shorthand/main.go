//shorthand example.
package main

import "fmt"

//declared example
// var b string     --- var name declaration

//initialize example
// b = "cow"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	// %v - value
	// %f - data type with value?
	// %T - data type

	fmt.Printf("%v \n", a)
	fmt.Printf("%f \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%f \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%f \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%f \n", d)

}