/*Exercise 2 — Typed constants

Create:

basics/constants/exercise-02-typed-constant/main.go

Create:

PASSMARK → 50
APPNAME  → "Go Learning Lab"

Requirements:

PASSMARK must explicitly be int
APPNAME must explicitly be string

Print both.

Expected:

50
Go Learning Lab*/

package main

import "fmt"


const PASSMARK int = 50
const APPNAME string = "Go Learning Lab"

func main() {
	fmt.Println(PASSMARK)
	fmt.Println(APPNAME)
}