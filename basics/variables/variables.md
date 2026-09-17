First understand what a variable is

A variable is simply a named place where a program stores a value.

Think of it like a labeled box:

┌───────────────┐
│ myNum         │
│               │
│ 50            │
└───────────────┘

The label is:

myNum

and the value stored inside is:

50

So when Go sees:

var myNum int = 50

you can read it as:

Create a variable called myNum, make it an integer, and store 50 inside it.

The pieces are:

var       myNum       int       =       50
 ↓          ↓          ↓        ↓        ↓
declare    name       type    assign    value
The main variable types from this lesson

These four are especially important:

var age int = 27
var price float32 = 19.99
var name string = "Ridge"
var learning bool = true

An int stores whole numbers:

27
-5
100

A float32 stores decimal numbers:

19.99
3.14
-2.5

A string stores text:

"Hello"
"Ridge Junior"
"Learning Go"

A bool stores only:

true
false
Three ways you will often declare variables

The fully explicit version is:

var student string = "John"

Here you tell Go everything:

name = student
type = string
value = "John"

You can also let Go infer the type:

var student = "John"

Go looks at:

"John"

and determines:

this must be a string

Inside a function, you can use the short declaration:

student := "John"

This:

:=

means:

Create a new variable and assign a value immediately.

So:

age := 27

is roughly equivalent to:

var age int = 27

when used inside a function.

One distinction you should memorize now:

:=  creates a new variable and assigns it
=   assigns a value to an already-declared variable

For example:

var age int
age = 27

The first line creates age.

The second line puts 27 inside it.

But:

age := 27

does both in one line.