package main

import "fmt"

func main(){
	fmt.Printf("Please enter your name: ")

	var name string
	score := 0
	num_questions := 2

	fmt.Scan(&name)

	fmt.Printf("Hello, %v, Welcome to the game\n", name)

	var age uint
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You are old enough to play")
	} else {
		fmt.Println("Comback later when your old enough")
	}
	var answer string	
	fmt.Printf("What fruit is better, orange or pineapple? ")
	fmt.Scan(&answer)

	if answer == "pineapple" {
		fmt.Println("Correct")
		score++
	}else {
		fmt.Println("Sorry not correct :(")
	}

	fmt.Printf("How many star wars movies are there? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 6 {
		fmt.Println("Correct")
	}else {
		fmt.Println("Sorry not correct :(")
	}
	
	fmt.Printf("You scored %v out of %v", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}