package main

import "fmt"

func main(){
	fmt.Printf("Please enter your name: ")

	var name string
	var age int

	fmt.Scan(&name)

	fmt.Println("Hello, %v, Welcome to the game",name )

	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println(age >=18)
}