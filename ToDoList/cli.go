package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"todolist/store"
)

type CommandArguments struct {
	scanner   *bufio.Scanner
	itemStore *store.ItemStore
}

var commands = map[string]func(args *CommandArguments) error{
	"a": cmdAdd,
	"r": cmdRemove,
	"l": cmdList,
}

func getDefaultDataPath() (string, error) {
	defaultDataPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	defaultDataPath = filepath.Join(defaultDataPath, "TodoList.json")
	return defaultDataPath, nil
}

func run() error {
	itemStore, err := loadStore()
	if err != nil {
		return err
	}
	err = loop(itemStore)

	return err
}

func loadStore() (*store.ItemStore, error) {
	dataPath, err := getDefaultDataPath()
	if err != nil {
		fmt.Printf("Got an error trying to fetch the default path %q\n", err)
		return nil, err
	}
	itemStore, err := store.ReadDB(dataPath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Database empty, building new DB...")
			itemStore = store.NewItemStore(dataPath)
		} else {
			err = fmt.Errorf("Unexpected error reading from DB: %w", err)
			return nil, err
		}
	}

	return itemStore, nil
}

func loop(itemStore *store.ItemStore) error {
	sc := bufio.NewScanner(os.Stdin)
	for {
		input, err := fetch("Enter command (a=add, r=remove, l=list, e=exit): ", sc)
		if err != nil {
			return err
		}
		if input == "e" {
			return nil
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
		err = command(arguments)
		if err != nil {
			fmt.Printf("There was an error running the command:", err)
		}
	}
}

func fetch(prompt string, sc *bufio.Scanner) (string, error) {
	fmt.Print(prompt)
	ok := sc.Scan()
	if ok {
		return sc.Text(), nil
	}
	return "", io.EOF
}

func cmdAdd(args *CommandArguments) error {
	name, err := fetch("Enter the name of the new entry: ", args.scanner)
	if err != nil {
		fmt.Printf("Error adding %q: %v\n", name, err)
		return err
	}

	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Printf("No empty name allowed\n")
		return nil
	}

	_, found := args.itemStore.SearchEntryByName(name)
	if found {
		fmt.Printf("Entry with name %q already exists.\n", name)
		return nil
	}
	entry := store.CreateEntry(name)
	args.itemStore.AddEntry(entry)
	tries := 3
	for range tries {
		err = args.itemStore.Save()
		if err != nil {
			fmt.Println("Error storing changes to the DB", err)
		} else {
			break
		}
	}
	return err
}

func cmdRemove(args *CommandArguments) error {
	args.itemStore.PrintEntries()
	name, err := fetch("Enter the name of the entry to remove: ", args.scanner)

	if err != nil {
		fmt.Printf("Error removing %q: %v\n", name, err)
		return err
	}

	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("No name input")
		return nil
	}

	err = args.itemStore.RemoveEntryByName(name)
	if err != nil {
		fmt.Printf("Error removing %q: %v\n", name, err)
		return err
	}

	fmt.Printf("Entry '%s' removed successfully.\n", name)
	err = args.itemStore.Save()

	if err != nil {
		fmt.Println("Error storing changes to the DB", err)
		return err
	}
	return nil
}

func cmdList(args *CommandArguments) error {
	args.itemStore.PrintEntries()
	return nil
}
