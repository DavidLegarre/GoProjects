package main

import (
	"fmt"
	"todolist/types"
)

var sampleEntry types.TodoEntry = types.TodoEntry{Name: "Sample Entry"}
var Entries types.TodoList = types.TodoList{Entries: []types.TodoEntry{sampleEntry}}

func createEntry(name string) types.TodoEntry {
	return types.TodoEntry{Name: name}
}

func AddEntry(entry types.TodoEntry) {
	Entries.Entries = append(Entries.Entries, entry)
}

func RemoveEntry(entry types.TodoEntry) {
	for i, e := range Entries.Entries {
		if e == entry {
			Entries.Entries = append(Entries.Entries[:i], Entries.Entries[i+1:]...)
			break
		}
	}
}

func GetEntries() []types.TodoEntry {
	return Entries.Entries
}

func ListEntries() {
	fmt.Println("Current To-Do List:")
	fmt.Println()
	for _, entry := range Entries.Entries {
		fmt.Println(entry.Name)
	}
	fmt.Println()
}
