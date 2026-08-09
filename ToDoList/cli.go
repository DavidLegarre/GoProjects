package main

import (
	"bufio"
	"errors"
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

func run() error {
	itemStore, err := loadStore()
	if err != nil {
		return err
	}
	loop(itemStore)

	return nil
}

func loadStore() (*store.ItemStore, error) {
	itemStore, err := store.ReadDB()

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Database empty, building new DB...")
			itemStore = store.NewItemStore()
		} else {
			fmt.Printf("Unexpected error occured %s", err)
			fmt.Println("Exiting...")
			return nil, err
		}
	}

	return itemStore, nil
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
	err := args.itemStore.Save()
	if err != nil {
		fmt.Println("Error storing changes to the DB", err)
	}
}

func cmdRemove(args *CommandArguments) {
	args.itemStore.PrintEntries()
	name := fetch("Enter the name of the entry to remove: ", args.scanner)

	err := args.itemStore.RemoveEntryByName(name)
	if err != nil {
		fmt.Printf("Error removing %q: %v\n", name, err)
		return
	}
	fmt.Printf("Entry '%s' removed successfully.\n", name)
	err = args.itemStore.Save()
	if err != nil {
		fmt.Println("Error storing changes to the DB", err)
	}

}

func cmdList(args *CommandArguments) {
	args.itemStore.PrintEntries()
}
