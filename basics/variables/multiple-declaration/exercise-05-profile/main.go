/*Exercise 5 — Mini profile

Create:

exercise-05-profile/main.go

Use a combination of multiple declaration and a variable block to store:

firstName = "Ridge"
lastName  = "Junior"
age       = 27
country   = "Kenya"
learning  = true

Rules:

firstName and lastName
→ declare together

age and country
→ declare in a var block

learning
→ declare using :=

Then print everything.

The key lesson to remember is:

var a, b int = 1, 2
→ same explicit type

var a, b = 1, "Hello"
→ Go infers different types

a, b := 1, "Hello"
→ short declaration inside a function

var (
    ...
)
→ grouped declarations for readability*/


package main

import "fmt"

var age, country = 27, "Kenya"

func main() {
	firstName, lastName := "Ridge", "Junior"
	learning := true
	fmt.Println(firstName, lastName, age, country, learning)
}