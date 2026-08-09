package main

import (
	"bufio"
	"fmt"
	"os"
	"todolist/store"
)

type CommandArguments struct {
	scanner   *bufio.Scanner
	itemStore *store.ItemStore
}

var commands = map[string]func(*CommandArguments){
	"a": cmdAdd,
	"r": cmdRemove,
	"l": cmdList,
}

func run() {

	itemStore := store.ReadDB()
	if itemStore == nil {
		fmt.Println("Failed to read the database. Exiting.")
		return
	}
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
		arguments := &CommandArguments{
			scanner:   sc,
			itemStore: itemStore,
		}
		command(arguments)
	}
}

func fetch(prompt string, sc *bufio.Scanner) string {
	fmt.Print(prompt)
	sc.Scan()
	return sc.Text()
}

func cmdAdd(args *CommandArguments) {
	entry := store.CreateEntry(fetch("Enter the name of the new entry: ", args.scanner))
	args.itemStore.AddEntry(entry)
	store.WriteDB()
}

func cmdRemove(args *CommandArguments) {
	args.itemStore.GetEntries()
	name := fetch("Enter the name of the entry to remove: ", args.scanner)
	entry := args.itemStore.SearchEntryByName(name)
	if entry != nil {
		args.itemStore.RemoveEntry(entry)
		store.WriteDB()
		fmt.Printf("Entry '%s' removed\n", name)
	} else {
		fmt.Println("Entry not found.")
	}
}

func cmdList(args *CommandArguments) {
	args.itemStore.PrintEntries()
}
