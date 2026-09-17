package main

import "fmt"

const (
	APPNAME       = "Go Learning Config Dashboard"
	VERSION       = "1.0.0"
	COUNTRY       = "Kenya"
	PLATFORM      = "W3Schools"
	COURSE        = "Go Programming"
	MAX_SCORE int = 100
	PASSMARK  int = 50
)

func main() {
	const CURRENTLEVEL string = "Beginner"
	const CURRENT_YEAR int = 2026

	var firstName, lastName string = "John", "Doe"
	city := "Nairobi"

	var currentScore, previousScore int = 75, 45
	isLearningActive := true

	var certificateEarned bool
	var bonusPoints int

	fmt.Println("===== GO LEARNING CONFIG DASHBOARD =====")
	fmt.Println()

	fmt.Println("Application:", APPNAME)
	fmt.Println("Version:", VERSION)
	fmt.Println("Country:", COUNTRY)
	fmt.Println("Platform:", PLATFORM)

	fmt.Println("================================================")

	fmt.Println("Course:", COURSE)
	fmt.Println("Current Level:", CURRENTLEVEL)
	fmt.Println("Current Year:", CURRENT_YEAR)
	fmt.Println("Maximum Score:", MAX_SCORE)
	fmt.Println("Pass Mark:", PASSMARK)

	fmt.Println("================================================")

	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("City:", city)
	fmt.Println("Previous Score:", previousScore)
	fmt.Println("Current Score:", currentScore)

	currentScore = 85
	fmt.Println("Updated Score:", currentScore)

	fmt.Println("================================================")

	fmt.Println("Is Learning Active:", isLearningActive)
	fmt.Println("Certificate Earned:", certificateEarned)
	fmt.Println("Bonus Points:", bonusPoints)
}