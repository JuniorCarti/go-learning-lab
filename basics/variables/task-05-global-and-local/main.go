/*Task 5 — Global and Local Variables

This one is slightly harder.

Create:

basics/variables/task-05-global-and-local/main.go

You need one variable outside main():

country → "Kenya"

Then inside main() create:

city → "Kisumu"
year → 2026

Use var for the variable outside the function.

Use := for the variables inside main().

Your structure should begin roughly like this:

package main

import "fmt"

// create country here

func main() {
	// create city and year here

	// print everything
}

Expected output:

Kenya
Kisumu
2026

The important lesson is:

var
→ can be used outside functions

:=
→ only inside functions*/

package main

import "fmt"

var country string = "Kenya"

func main() {
	city := "Kisumu"
	year := 2026
	fmt.Println(country)
	fmt.Println(city)
	fmt.Println(year)
}