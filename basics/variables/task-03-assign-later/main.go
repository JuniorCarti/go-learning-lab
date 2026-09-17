/*Task 3 — Declare Now, Assign Later

Create:

basics/variables/task-03-assign-later/main.go

Start with:

package main

import "fmt"

func main() {
	var language string
	var year int
	var learning bool

	// Assign values here

	fmt.Println(language)
	fmt.Println(year)
	fmt.Println(learning)
}

Assign:

language → "Go"
year     → 2026
learning → true

Important rule:

Use:

=

for the assignments.

Do not redeclare the variables with :=.

Expected output:

Go
2026
true*/


package main

import "fmt"

func main() {
	var language string
	var year int
	var learning bool

	language = "Go"
	year = 2026
	learning = true

	fmt.Println(language)
	fmt.Println(year)
	fmt.Println(learning)
}