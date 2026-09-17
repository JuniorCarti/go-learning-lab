/*write a program 
Write a program containing these three variables:

name     → "Ridge"
age      → 27
learning → true

Use:

var

for name, but use:

:=

for age and learning.

Then print all three.

Your expected output should be:

Ridge
27
true*/


package main

import "fmt"

var name = "Ridge"

func main() {
	age := 27
	learning := true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(learning)
}
