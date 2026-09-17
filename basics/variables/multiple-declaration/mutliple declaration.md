The main idea is simple:

Go lets you create several variables at once instead of writing one declaration per line.

For example, instead of:

var a int = 1
var b int = 3
var c int = 5
var d int = 7

you can write:

var a, b, c, d int = 1, 3, 5, 7

Think of the values matching the variables by position:

a    b    c    d
↓    ↓    ↓    ↓
1    3    5    7

So:

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)

prints:

1
3
5
7
Same type on one line

This works:

var a, b, c int = 10, 20, 30

because all three variables are int.

You can also let Go infer the type:

var a, b, c = 10, 20, 30

Go sees whole numbers and understands they are integers.

But there is an important rule from your material: when you explicitly give one type after the variable names, that declaration is for that one type.

So this would not make sense:

var a, b int = 10, "Hello"

because "Hello" is a string, not an int.

Mixed types on one line

If you leave out the explicit type, Go can infer different types:

var a, b = 6, "Hello"

Now:

a → int
b → string

Go figures that out from the values.

You can do the same using := inside a function:

c, d := 7, "World!"

Here:

c → int
d → string

So this is valid:

package main

import "fmt"

func main() {
	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

Output:

6
Hello
7
World!
Variable blocks

Go also lets you group variable declarations:

var (
	a int
	b int = 1
	c string = "hello"
)

This is useful when you have several related variables and want the code to be easier to read.

You can think of:

var (
	...
)

as:

“I am about to declare a group of variables.”

This:

var (
	name string = "Ridge"
	age  int    = 27
	city string = "Kisumu"
)

is often cleaner than:

var name string = "Ridge"
var age int = 27
var city string = "Kisumu"

Both are valid.

One detail from your lesson is especially important:

var a, b = 6, "Hello"

can infer different types.

But:

var a, b int = 6, 7

uses one explicit type for both.

And:

a, b := 6, "Hello"

also allows different inferred types, but := must be used inside a function.