package main

import (
	"encoding/json"
	"fmt"
	"slices"
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

func GetEntries() []types.TodoEntry {
	return Items.Entries
}

func PrintEntries() {
	fmt.Println("Current To-Do List:")
	fmt.Println()
	for _, entry := range Items.Entries {
		fmt.Println(entry.Name)
	}
	fmt.Println()
}

func ReadDB() {
	data := readJsonFile()
	if data == nil {
		fmt.Println("Error reading JSON file.")
		return
	}

	var todoList types.TodoList
	err := json.Unmarshal(data, &todoList)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	Items = todoList
}

func WriteDB() {
	b, err := json.Marshal(Items)
	if err != nil {
		fmt.Println("Error marshalling list:", err)
		return
	}
	writeJsonFile(b)
}

func DeleteDB() {
	Items = types.TodoList{Entries: []types.TodoEntry{}}
	WriteDB()
}

func RemoveEntryByName(name string) {
	for i, entry := range Items.Entries {
		if entry.Name == name {
			Items.Entries = slices.Delete(Items.Entries, i, i+1)
			return
		}
	}
}
