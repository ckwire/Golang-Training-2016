package main

import "fmt"

const p string = "death & taxes"

const (
	Pi = 3.14
	Language = "Go"
)

func main() {

	const q = 42

	fmt.Println("p -  ", p)
	fmt.Println("q -  ", q)
	fmt.Println("PI -  ", Pi)
	fmt.Println("Language -  ", Language)

}