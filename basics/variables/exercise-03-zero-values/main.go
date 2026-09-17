//exercise is to create zero values 

/*package main

import "fmt"

func main() {
	var age int
	var name string
	var d bool

	fmt.Println(age, name, d)
*/


/*Go automatically gives variables default values.

A string gets:

""

an int gets:

0

and a bool gets:

false

So this exercise is teaching you that Go variables are initialized even if you don't explicitly give them a value.*/
package main

import "fmt"

func main() {
	var name string
	var age int
	var active bool

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(active)
}