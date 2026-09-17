package main

import "fmt"

const (
	// This is a constant declaration block. You can define multiple constants here.	
	APPNAME    = "Go Learning Config Dashboard"
	VERSION    = "1.0.0"
	COUNTRY    = "Kenya"
	PLATFORM  = "W3Schools"
	COURSE	  = "Go Programming"
	MAX_SCORE int  = 100
	PASSMARK int = 50

)

func main() {

	const CURRENTLEVEL string = "Beginner"
	const CURRENT_YEAR int = 2024
	var firstName, lastName string = "John", "Doe"
	city := "Nairobi"
	var currentScore, previousScore int = 75, 45
	isLearningActive := true
	var certificateEarned bool
	var bonusPoints int

	fmt.Println("Welcome to the", APPNAME)

	fmt.Println("================================================")
	fmt.Println("Version:", VERSION)
	fmt.Println("Country:", COUNTRY)
	fmt.Println("Platform:", PLATFORM)

	fmt.Println("================================================")
	fmt.Println("Course:", COURSE)
	fmt.Println("Current Level:", CURRENTLEVEL)
	fmt.Println("Current Year:", CURRENT_YEAR)
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("City:", city)
	fmt.Println("Current Score:", currentScore)
	fmt.Println("Previous Score:", previousScore)

	fmt.Println("==================================================")
	fmt.Println("Is Learning Active:", isLearningActive)
	fmt.Println("Certificate Earned:", certificateEarned)
	fmt.Println("Bonus Points:", bonusPoints)

}