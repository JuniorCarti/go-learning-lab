/*Exercise 4 — Constant block

Create:

basics/constants/exercise-04-constant-block/main.go

Use one:

const (
)

block to create:

COURSE   = "Go Programming"
PLATFORM = "W3Schools"
MAXSCORE = 100
ACTIVE   = true

Print all four.

Expected:

Go Programming
W3Schools
100
true*/

package main

import "fmt"

const (
	COURSE   = "Go Programming"
	PLATFORM = "W3Schools"
	MAXSCORE = 100
	ACTIVE   = true
)

func main() {
	fmt.Println(COURSE)
	fmt.Println(PLATFORM)
	fmt.Println(MAXSCORE)
	fmt.Println(ACTIVE)
}
