package store

import (
	"fmt"
	"slices"
	"todolist/types"
)

type ItemStore struct {
	items types.TodoList
}

func NewItemStore() *ItemStore {
	return &ItemStore{items: types.TodoList{Entries: nil}}
}

func CreateEntry(name string) types.TodoEntry {
	return types.TodoEntry{Name: name}
}

func (s *ItemStore) AddEntry(entry types.TodoEntry) {
	s.items.Entries = append(s.items.Entries, entry)
}

func (s *ItemStore) RemoveEntryByName(name string) bool {
	for i, e := range s.items.Entries {
		if e.Name == name {
			s.items.Entries = slices.Delete(s.items.Entries, i, i+1)
			return true
		}
	}
	fmt.Printf("Entry with name '%s' not found.\n", name)

	return false
}

func (s *ItemStore) GetEntries() types.TodoList {
	return s.items
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

func (s *ItemStore) Save() {
	WriteDB(s)
}
