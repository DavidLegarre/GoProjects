package store

import (
	"errors"
	"fmt"
	"slices"
	"todolist/types"
)

var ErrNotFound = errors.New("entry not found")

type ItemStore struct {
	items types.TodoList
	path  string
}

func NewItemStore(path string) *ItemStore {
	return &ItemStore{items: types.TodoList{Entries: nil}, path: path}
}

func CreateEntry(name string) types.TodoEntry {
	return types.TodoEntry{Name: name}
}

func (s *ItemStore) AddEntry(entry types.TodoEntry) {
	s.items.Entries = append(s.items.Entries, entry)
}

func (s *ItemStore) RemoveEntryByName(name string) error {
	ogLen := len(s.items.Entries)
	s.items.Entries = slices.DeleteFunc(s.items.Entries, func(entry types.TodoEntry) bool {
		return entry.Name == name
	})
	if len(s.items.Entries) != ogLen {
		return nil
	} else {
		return ErrNotFound
	}
}

func (s *ItemStore) GetEntries() types.TodoList {
	return s.items
}

func (s *ItemStore) SearchEntryByName(name string) *types.TodoEntry {
	for i := range s.items.Entries {
		if s.items.Entries[i].Name == name {
			return &s.items.Entries[i]
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

func (s *ItemStore) Save() error {
	return WriteDB(s)
}
