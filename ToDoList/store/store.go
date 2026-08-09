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

func (s *ItemStore) RemoveEntryByName(name string) error {
	for i, e := range s.items.Entries {
		if e.Name == name {
			s.items.Entries = slices.Delete(s.items.Entries, i, i+1)
			return nil
		}
	}
	return ErrNotFound
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
