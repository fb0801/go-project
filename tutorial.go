package main

import "fmt"

func main(){
	fmt.Printf("Please enter your name: ")

	var name string
	

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
	}else {
		fmt.Println("Sorry not correct :(")
	}
}