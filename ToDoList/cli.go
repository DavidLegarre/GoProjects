package main

import "fmt"

func Run() {
	for {
		fmt.Println("Hello Loop!")
		input := fetchInput()
		switch input {
		case "a":
			fmt.Println("Adding a new entry...")
		case "e":
			fmt.Println("Exiting...")
			return
		}
	}
}

func fetchInput() (input string) {
	fmt.Printf("Enter your input: ")
	fmt.Scanln(&input)
	return input
}
