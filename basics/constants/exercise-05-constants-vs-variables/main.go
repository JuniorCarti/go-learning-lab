/*Exercise 5 — Variables and constants together

Create:

basics/constants/exercise-05-constants-vs-variables/main.go

Create:

MAXSCORE = 100
PASSMARK = 50

as constants.

Create:

studentName  = "Ridge"
studentScore = 75

as variables.

Print everything.

Then update:

studentScore

from:

75

to:

85

and print it again.

Do not change:

MAXSCORE
PASSMARK

Expected output could be:

Student: Ridge
Maximum Score: 100
Pass Mark: 50
Original Score: 75
Updated Score: 85

This is the most important exercise because it makes the distinction clear:

studentScore
→ data that changes

MAXSCORE
→ fixed rule8*/

package main

import "fmt"

const (
	MAXSCORE = 100
	PASSMARK = 50
)

var (
	studentName  = "Ridge"
	studentScore = 75
)
func main() {
	fmt.Println("Student:", studentName)
	fmt.Println("Maximum Score:", MAXSCORE)
	fmt.Println("Pass Mark:", PASSMARK)
	fmt.Println("Original Score:", studentScore)
	studentScore = 85
	fmt.Println("Updated Score:", studentScore)
}