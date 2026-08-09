package main

import (
	"bufio"
	"fmt"
	"os"
)

var commands = map[string]func(*bufio.Scanner){
	"a": cmdAdd,
	"r": cmdRemove,
	"l": cmdList,
}

func run() {

	ReadDB()
	sc := bufio.NewScanner(os.Stdin)
	for {
		input := fetch("Enter command (a=add, r=remove, l=list, e=exit): ", sc)
		if input == "e" {
			return
		}

		command, ok := commands[input]
		if !ok {
			fmt.Println("Invalid command. Please try again.")
			continue
		}
		command(sc)
	}
}

func fetch(prompt string, sc *bufio.Scanner) string {
	fmt.Print(prompt)
	sc.Scan()
	return sc.Text()
}

func cmdAdd(sc *bufio.Scanner) {
	entry := createEntry(fetch("Enter the name of the new entry: ", sc))
	AddEntry(entry)
	WriteDB()
}

func cmdRemove(sc *bufio.Scanner) {
	PrintEntries()
	name := fetch("Enter the name of the entry to remove: ", sc)
	RemoveEntryByName(name)
	WriteDB()
}

func cmdList(sc *bufio.Scanner) {
	PrintEntries()
}
