package main

import "fmt"

var country = "Kenya"
var repository = "go-learning-lab"

func main() {
	var firstName, lastName string = "Ridge", "Junior"
	var exercisesCompleted, projectsCompleted int = 12, 1

	var age = 27
	var platform = "W3Schools"

	city := "Kisumu"
	course := "Go Programming"
	learning := true

	var (
		currentTopic string = "Multiple Variable Declaration"
		score        int    = 85
		targetScore  int    = 100
		language     string = "Go"
	)

	var certificateName string
	certificateName = "Go Developer Certificate"

	var accountVerified bool
	var certificateEarned bool
	var bonusPoints int

	fmt.Println("===== STUDENT LEARNING PROFILE =====")

	fmt.Println()
	fmt.Println("Country:", country)
	fmt.Println("Repository:", repository)

	fmt.Println("===================================")
	fmt.Println("Name:", firstName, lastName)
	fmt.Println("Exercises Completed:", exercisesCompleted)
	fmt.Println("Projects Completed:", projectsCompleted)
	fmt.Println("Age:", age)
	fmt.Println("Platform:", platform)
	fmt.Println("City:", city)

	fmt.Println("===================================")
	fmt.Println("Course:", course)
	fmt.Println("Learning:", learning)
	fmt.Println("Current Topic:", currentTopic)
	fmt.Println("Score:", score)
	fmt.Println("Target Score:", targetScore)
	fmt.Println("Language:", language)

	fmt.Println("===================================")
	fmt.Println("Certificate Name:", certificateName)
	fmt.Println("Account Verified:", accountVerified)
	fmt.Println("Certificate Earned:", certificateEarned)
	fmt.Println("Bonus Points:", bonusPoints)
}
