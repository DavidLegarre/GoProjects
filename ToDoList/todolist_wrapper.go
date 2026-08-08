package main

import (
	"encoding/json"
	"fmt"
	"todolist/types"
)

var sampleEntry types.TodoEntry = types.TodoEntry{Name: "Sample Entry"}
var Items types.TodoList = types.TodoList{Entries: []types.TodoEntry{sampleEntry}}

func createEntry(name string) types.TodoEntry {
	return types.TodoEntry{Name: name}
}

func AddEntry(entry types.TodoEntry) {
	Items.Entries = append(Items.Entries, entry)
}

func EntryToBytes(entry types.TodoEntry) []byte {
	b, err := json.Marshal(entry)

	if err != nil {
		fmt.Println("Error marshalling entry:", err)
		return nil
	}

	return b
}

func RemoveEntry(entry types.TodoEntry) {
	for i, e := range Items.Entries {
		if e == entry {
			Items.Entries = append(Items.Entries[:i], Items.Entries[i+1:]...)
			break
		}
	}
}

func GetEntries() []types.TodoEntry {
	return Items.Entries
}

func ListEntries() {
	fmt.Println("Current To-Do List:")
	fmt.Println()
	for _, entry := range Items.Entries {
		fmt.Println(entry.Name)
	}
	fmt.Println()
}
