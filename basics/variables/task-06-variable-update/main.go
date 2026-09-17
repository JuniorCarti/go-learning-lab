/*fter completing all five, create one extra program:

basics/variables/task-06-variable-update/main.go

Start with:

score := 10

Print it, change it to:

25

and print it again.

Expected:

10
25

The key question is whether the second statement should use:

=

or:

:=*/

package main

import "fmt"

func main() {
	score := 10
	fmt.Println(score)
	score = 25
	fmt.Println(score)
}
