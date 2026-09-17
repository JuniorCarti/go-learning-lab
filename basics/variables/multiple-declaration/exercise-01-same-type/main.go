/*Exercise 1 — Same type

Create:

basics/variables/multiple-declaration/exercise-01-same-type/main.go

Declare these four variables on one line:

a = 2
b = 4
c = 6
d = 8

All must explicitly be:

int

Print each one.

Expected:

2
4
6
8*/

package main

import "fmt"

var a, b, c, d int = 2, 4, 6, 8

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}