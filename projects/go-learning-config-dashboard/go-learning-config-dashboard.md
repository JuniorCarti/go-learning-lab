Major Project — Go Learning Configuration Dashboard

Store it here:

go-learning-lab/
└── projects/
    └── go-learning-config-dashboard/
        └── main.go

You're going to build a program representing the configuration and current state of your Go learning system.

This will combine what you've learned about:

variables
multiple-variable declarations
variable blocks
constants
typed constants
untyped constants
global values
local values
assignment
zero values
fmt.Println

Your finished output should look roughly like:

===== GO LEARNING DASHBOARD =====

Application: Go Learning Lab
Version: 1.0
Country: Kenya
Platform: W3Schools

Student: Ridge Junior
City: Kisumu

Course: Go Programming
Maximum Score: 100
Pass Mark: 50

Current Score: 78
Previous Score: 65
Learning Active: true

Certificate Earned: false
Bonus Points: 0

Your job is not just to make the output appear. You must decide which information should be constant and which should be variable.

Fixed configuration

Outside main(), create a constant block containing:

APPNAME   = "Go Learning Lab"
VERSION   = "1.0"
COUNTRY   = "Kenya"
PLATFORM  = "W3Schools"
COURSE    = "Go Programming"
MAXSCORE  = 100
PASSMARK  = 50

Use explicit types for at least:

MAXSCORE
PASSMARK

For example:

const (
	APPNAME = ...
	MAXSCORE int = ...
)

These are constants because they represent configuration that should not change while the program runs.

Inside main(), use variables for:

firstName
lastName
city
currentScore
previousScore
learningActive
certificateEarned
bonusPoints

Use different declaration techniques.

For example:

firstName and lastName must be declared together:

var firstName, lastName string = ...

city should use:

:=

currentScore and previousScore should be declared together:

var currentScore, previousScore int = ...

learningActive should use:

:=

And create:

var certificateEarned bool
var bonusPoints int

without assigning values so that you demonstrate zero values again.

Now update one variable.

Start with:

currentScore = 78

Then later:

currentScore = 85

Print both stages.

This proves you understand:

variable
→ can change

while your constants remain fixed.

Add one local constant inside main():

const CURRENTLEVEL = "Beginner"

That proves you understand that constants can exist both:

outside functions

and:

inside functions

Your project should therefore demonstrate this relationship:

SYSTEM CONFIGURATION
        │
        ├── APPNAME       constant
        ├── VERSION       constant
        ├── MAXSCORE      constant
        └── PASSMARK      constant

STUDENT STATE
        │
        ├── currentScore       variable
        ├── city               variable
        ├── learningActive     variable
        └── certificateEarned  variable

That distinction is extremely important in real software.

Configuration and business rules are often fixed for a particular run of the application, while user data and application state change.

Build the project yourself without copying a completed solution. When you've written main.go, send me the entire code and I'll review it the same way we reviewed your Student Learning Profile project—line by line, including whether each value is correctly modeled as a constant or variable