package main

import (
	"bufio"
	"fmt"
	"os"
)

func run() {
	for {
		sc := bufio.NewScanner(os.Stdin)
		input := fetch("Please enter a command (a: add, l: list, e: exit): ", sc)
		switch input {
		case "a":
			fmt.Println("Adding a new entry...")
			fmt.Println()
			entry := createEntry(fetch("Enter the name of the new entry: ", sc))
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

func fetch(prompt string, sc *bufio.Scanner) string {
	fmt.Print(prompt)
	sc.Scan()
	return sc.Text()
}
