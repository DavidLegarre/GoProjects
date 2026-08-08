package main

import (
	"fmt"
	"todolist/types"
)

var sampleEntry types.TodoEntry = types.TodoEntry{Name: "Sample Entry"}
var Entries types.TodoList = types.TodoList{Items: []types.TodoEntry{sampleEntry}}

func createEntry(name string) types.TodoEntry {
	return types.TodoEntry{Name: name}
}

func AddEntry(entry types.TodoEntry) {
	Entries.Items = append(Entries.Items, entry)
}

func RemoveEntry(entry types.TodoEntry) {
	for i, e := range Entries.Items {
		if e == entry {
			Entries.Items = append(Entries.Items[:i], Entries.Items[i+1:]...)
			break
		}
	}
}

func GetEntries() []types.TodoEntry {
	return Entries.Items
}

func ListEntries() {
	fmt.Println("Current To-Do List:")
	fmt.Println()
	for _, entry := range Entries.Items {
		fmt.Println(entry.Name)
	}
	fmt.Println()
}
