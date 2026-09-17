/*Exercise 4 — Variable block

Create:

exercise-04-variable-block/main.go

Use:

var (
	...
)

to create:

username = "Junior"
score    = 100
active   = true

Print all three.

Expected:

Junior
100
true*/

package main

import "fmt"

var (
	username = "Junior"
	score    = 100
	active   = true
)

func main() {
	fmt.Println(username)
	fmt.Println(score)
	fmt.Println(active)
}