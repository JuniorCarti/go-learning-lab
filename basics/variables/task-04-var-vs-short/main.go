/*Task 4 — var vs :=

Create:

basics/variables/task-04-var-vs-short/main.go

Create these variables:

course   → "Go Programming"
lessons  → 20
finished → false
price    → 19.99

But follow these rules:

course   → must use var
lessons  → must use :=
finished → must use :=
price    → must use var with an explicit type

Then print everything.

For example, part of your program should resemble:

var course = ...
lessons := ...

This task checks whether you understand that Go can infer types.*/

package main

import "fmt"

var course = "Go Programming"

func main() {
	lessons := 20
	finished := false
	var price float64 = 19.99

	fmt.Println(course)
	fmt.Println(lessons)
	fmt.Println(finished)
	fmt.Println(price)
}