package main

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	for {
		fmt.Println("Hello Loop!")
		input := fetchInput()
		switch input {
		case "a":
			fmt.Println("Adding a new entry...")
			fmt.Println()
			entry := createEntry(fetchName())
			AddEntry(entry)
		case "l":
			fmt.Println("Listing all entries...")
			ListEntries()
		case "e":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func fetchInput() (input string) {
	fmt.Printf("Enter your input: ")
	fmt.Scanln(&input)
	return input
}

func fetchName() (name string) {
	fmt.Printf("Enter the name of the entry: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
