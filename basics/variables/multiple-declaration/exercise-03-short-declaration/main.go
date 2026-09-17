/*Exercise 3 — Multiple :=

Create:

exercise-03-short-declaration/main.go

Inside main(), create these two variables in one statement:

language = "Go"
year     = 2026

You must use:

:=

Expected:

Go
2026*/

package main

import "fmt" 

func main() {
	language, year := "Go", 2026
	fmt.Println(language)
	fmt.Println(year)
}