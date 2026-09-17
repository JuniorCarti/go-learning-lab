/*Exercise 2 — Mixed types with var

Create:

exercise-02-mixed-types/main.go

Declare these on the same line:

age  = 27
name = "Ridge"

Do not explicitly write their types. Let Go infer them.

Then print both.

Expected:

27
Ridge*/

package main

import "fmt"

func main() {
	age, name := 27, "Ridge"
	fmt.Println(age)
	fmt.Println(name)
}