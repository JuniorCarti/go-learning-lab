1. What is a constant?

You already understand a variable as a labeled box:

var score = 85

You can later change it:

score = 90

A constant is different.

const PASSMARK = 50

means:

Create a value called PASSMARK, give it the value 50, and do not allow the program to change it later.

Think of the difference like this:

VARIABLE

score
┌───────────┐
│    85     │
└───────────┘
      ↓
can change
      ↓
┌───────────┐
│    90     │
└───────────┘

But:

CONSTANT

PASSMARK
┌───────────┐
│    50     │
└───────────┘

LOCKED 🔒

The program can read it, but cannot assign a new value to it.

2. Declaring a constant

The basic syntax is:

const NAME = value

For example:

const PI = 3.14

Then:

fmt.Println(PI)

prints:

3.14

You can also explicitly specify the type:

const PI float64 = 3.14

The general form is:

const      PI      float64      =      3.14
  ↓         ↓         ↓         ↓        ↓
keyword    name      type     assign    value
3. Constants must have a value immediately

With variables, this is valid:

var score int

Go gives it the zero value:

0

Then later:

score = 85

But you cannot do that with a constant.

This is invalid:

const SCORE int

because Go needs to know the fixed value when the constant is declared.

You must do:

const SCORE int = 85

So remember:

variable
→ can be created without a value

constant
→ must receive its value immediately
4. Constants cannot be reassigned

This works:

var score = 85
score = 90

But this does not:

const SCORE = 85
SCORE = 90

Go will give an error similar to:

cannot assign to SCORE

Why?

Because:

const SCORE = 85

means that 85 is supposed to remain fixed.

This is one of the biggest differences between:

var

and:

const
5. When should something be a constant?

Ask yourself:

Should this value change while my program is running?

If yes, use a variable.

For example:

var score = 85

A student's score may later become:

90

so score makes sense as a variable.

But suppose your system says:

const MAXSCORE = 100

The maximum possible score is fixed.

Or:

const PASSMARK = 50

If the pass mark is defined by the system and not supposed to change while the program runs, a constant makes sense.

Other examples:

const COUNTRY = "Kenya"
const APPNAME = "Go Learning Lab"
const VERSION = "1.0"
const DAYSINWEEK = 7
6. Typed constants

A typed constant explicitly specifies the type.

Example:

const A int = 1

Here we are explicitly telling Go:

A
↓
int
↓
1

Other examples:

const PASSMARK int = 50
const APPNAME string = "Go Learning Lab"
const ACTIVE bool = true
const RATE float64 = 1.5

You decide the type yourself.

7. Untyped constants

You can also leave the type out:

const A = 1

Go understands from the value that it is numeric.

Similarly:

const NAME = "Ridge"
const ACTIVE = true
const PI = 3.14

You don't explicitly write:

string
bool
float64

Go handles the constant according to its value and the context in which it is used.

For your current level, you can think of it as similar to type inference:

const NAME = "Ridge"

Go knows this represents a string constant.

Compare:

const AGE int = 27

with:

const AGE = 27

Both are valid.

The first is typed.

The second is untyped.

8. Why do untyped constants exist?

This is one place where Go constants are a little more powerful than ordinary variables.

Suppose:

const NUMBER = 10

Because it is untyped, Go can use that constant in compatible numeric contexts more flexibly.

At your current stage, you don't need to go deeply into this yet.

Just remember:

typed constant
→ you explicitly specify a type

untyped constant
→ Go leaves the exact type flexible until needed

You'll understand why this matters much more when we reach Go's numeric types.

9. Constants inside functions

You can declare constants inside:

func main()

For example:

package main

import "fmt"

func main() {
	const PASSMARK = 50

	fmt.Println(PASSMARK)
}

This constant belongs to that function's scope.

10. Constants outside functions

You can also declare them outside:

package main

import "fmt"

const PASSMARK = 50

func main() {
	fmt.Println(PASSMARK)
}

Now it is a package-level constant.

This is similar to what you learned with:

var country = "Kenya"

outside main().

So both:

var

and:

const

can be used outside functions.

But:

:=

cannot.

Compare:

package main

const PASSMARK = 50
var country = "Kenya"

// This would be invalid:
// age := 27

func main() {
	age := 27
}
11. Multiple constants

Just like variables, constants can be grouped.

Instead of:

const APPNAME = "Go Learning Lab"
const VERSION = "1.0"
const MAXSCORE = 100

you can write:

const (
	APPNAME  = "Go Learning Lab"
	VERSION  = "1.0"
	MAXSCORE = 100
)

This is called a constant block.

It's especially useful when the constants belong together.

For example:

const (
	COURSE     = "Go Programming"
	PLATFORM   = "W3Schools"
	PASSMARK   = 50
	MAXSCORE   = 100
	CERTIFIED  = false
)
12. Typed and untyped constants can coexist in a block

You can write:

const (
	MAXSCORE int = 100
	PI           = 3.14
	GREETING     = "Hello"
)

Here:

MAXSCORE
→ typed constant

PI
→ untyped constant

GREETING
→ untyped constant

That's completely valid.

13. Naming constants

The material you're using shows constants in uppercase:

const PI = 3.14
const A = 1

Uppercase names can make constants visually easy to distinguish while you're learning.

But Go itself does not require constant names to be uppercase.

These are both valid:

const PI = 3.14

and:

const Pi = 3.14

and even:

const pi = 3.14

In real Go projects, naming generally follows normal Go naming conventions rather than requiring every constant to use all capitals.

For this learning repo, uppercase names like:

PASSMARK
MAXSCORE
APPNAME

are perfectly fine because they make the idea obvious.

14. Constant versus variable

This distinction should become automatic:

Variable
──────────────
var score = 85

can become:

score = 90

versus:

Constant
──────────────
const MAXSCORE = 100

cannot become:

MAXSCORE = 200

A useful real-world example:

const MAXSCORE = 100

var studentScore = 75

MAXSCORE represents a system rule.

studentScore represents changing data.

Later:

studentScore = 90

is fine.

But:

MAXSCORE = 200

is not.

15. What values can constants contain?

At your level, focus on constants such as:

const AGE = 27
const PI = 3.14
const NAME = "Go Learning Lab"
const ACTIVE = true

That gives you:

integer constants
floating-point constants
string constants
boolean constants

Those are the main categories you'll encounter now.