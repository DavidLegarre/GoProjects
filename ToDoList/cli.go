package main

import (
	"bufio"
	"fmt"
	"os"
	"todolist/store"
)

var commands = map[string]func(*bufio.Scanner){
	"a": cmdAdd,
	"r": cmdRemove,
	"l": cmdList,
}

func run() {

	store.ReadDB()
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
	entry := store.CreateEntry(fetch("Enter the name of the new entry: ", sc))
	store.AddEntry(entry)
	store.WriteDB()
}

func cmdRemove(sc *bufio.Scanner) {
	store.PrintEntries()
	name := fetch("Enter the name of the entry to remove: ", sc)
	entry := store.SearchEntryByName(name)
	if entry != nil {
		store.RemoveEntry(entry)
		store.WriteDB()
		fmt.Printf("Entry '%s' removed\n", name)
	} else {
		fmt.Println("Entry not found.")
	}
}

func cmdList(_ *bufio.Scanner) {
	store.PrintEntries()
}
