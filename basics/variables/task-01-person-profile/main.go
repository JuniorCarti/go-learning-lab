/*Task 1 — Personal Profile

Create:

basics/variables/task-01-person-profile/main.go

Create four variables:

name       → string
age        → int
height     → float32
isLearning → bool

Give them values of your choice, then print all four.

Your output could look like:

Ridge
27
1.75
true

Requirements:

var name string = ...
var age int = ...
var height float32 = ...
var isLearning bool = ...

This task practices explicit variable types.*/

package main

import "fmt"

var name string = "Ridge"
var age int = 27
var height float32 = 1.75
var isLearning bool = true

func main() {
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(isLearning)
}
