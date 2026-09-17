package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = 3.14
	var s string = "Go"
	var b bool = true

	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", b, b)
}
