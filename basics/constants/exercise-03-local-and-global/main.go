/*Exercise 3 — Local and package constants

Create:

basics/constants/exercise-03-local-and-global/main.go

Outside main(), create:

COUNTRY = "Kenya"

Inside main(), create:

CITY = "Kisumu"

Then print both.

Expected:

Kenya
Kisumu

This tests your understanding of scope.*/

package main

import "fmt"

const COUNTRY = "Kenya"

func main() {
	const CITY = "Kisumu"
	fmt.Println(COUNTRY)
	fmt.Println(CITY)
}