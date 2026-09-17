/*Task 2 — Zero Values

Create:

basics/variables/task-02-zero-values/main.go

Declare these variables but do not assign values:

var username string
var score int
var loggedIn bool
var balance float32

Print all four.

Before running the program, write down what you think the output will be.

This task tests whether you understand Go's default values:

string  → ""
int     → 0
bool    → false
float32 → 0*/

package main

import "fmt"

var username string
var score int
var loggedIn bool
var balance float32

func main() {
	fmt.Println(username)
	fmt.Println(score)
	fmt.Println(loggedIn)
	fmt.Println(balance)
}