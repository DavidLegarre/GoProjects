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
	itemStore := loadStore()
	loop(itemStore)
}

func loadStore() *store.ItemStore {
	itemStore, err := store.ReadDB()

	if err == nil {
		fmt.Println("Database empty, building new DB...")
		itemStore = store.NewItemStore()
	}

	return itemStore
}

func loop(itemStore *store.ItemStore) {
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
	args.itemStore.Save()
}

func cmdRemove(args *CommandArguments) {
	args.itemStore.PrintEntries()
	name := fetch("Enter the name of the entry to remove: ", args.scanner)
	if args.itemStore.RemoveEntryByName(name) {
		fmt.Printf("Entry '%s' removed successfully.\n", name)
		args.itemStore.Save()
	}
}

func cmdList(args *CommandArguments) {
	args.itemStore.PrintEntries()
}
