package store

import (
	"encoding/json"
	"fmt"
	"slices"
	"todolist/types"
)

var sampleEntry types.TodoEntry = types.TodoEntry{Name: "Sample Entry"}
var Items types.TodoList = types.TodoList{Entries: []types.TodoEntry{sampleEntry}}

type ItemStore struct {
	items types.TodoList
}

func CreateEntry(name string) types.TodoEntry {
	return types.TodoEntry{Name: name}
}

func (s *ItemStore) AddEntry(entry types.TodoEntry) {
	s.items.Entries = append(s.items.Entries, entry)
}

func (s *ItemStore) RemoveEntry(entry *types.TodoEntry) {
	for i, e := range s.items.Entries {
		if e.Name == entry.Name {
			s.items.Entries = slices.Delete(s.items.Entries, i, i+1)
			return
		}
	}
}

func (s *ItemStore) GetEntries() []types.TodoEntry {
	return s.items.Entries
}

func (s *ItemStore) SearchEntryByName(name string) *types.TodoEntry {
	for _, entry := range s.items.Entries {
		if entry.Name == name {
			return &entry
		}
	}
	return nil
}

func (s *ItemStore) PrintEntries() {
	fmt.Println("Current To-Do List:")
	fmt.Println()
	for _, entry := range s.items.Entries {
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

func SearchEntryByName(name string) *types.TodoEntry {
	for _, entry := range Items.Entries {
		if entry.Name == name {
			return &entry
		}
	}
	return nil
}
