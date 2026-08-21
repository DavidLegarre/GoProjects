package store

import (
	"errors"
	"fmt"
	"sync"
	"todolist/types"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("entry not found")

type ItemStore struct {
	mu    sync.RWMutex
	items types.TodoList
	path  string
}

func NewItemStore(path string) *ItemStore {
	return &ItemStore{items: types.TodoList{EntryMap: make(map[uuid.UUID]types.TodoEntry)}, path: path}
}

func generateID() uuid.UUID {
	return uuid.New()
}

func CreateEntry(name string) types.TodoEntry {
	return types.TodoEntry{Id: generateID(), Name: name, DoneStatus: false}
}

func (s *ItemStore) AddEntry(entry types.TodoEntry) {
	s.items.EntryMap[entry.Id] = entry
}

func (s *ItemStore) PopEntry(id uuid.UUID) {
	delete(s.items.EntryMap, id)
}

func (s *ItemStore) RemoveEntryByName(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	found := false
	for id, entry := range s.items.EntryMap {
		if entry.Name == name {
			delete(s.items.EntryMap, id)
			found = true
		}
	}
	if !found {
		return ErrNotFound
	}
	return nil
}

func (s *ItemStore) SearchEntryByName(name string) (types.TodoEntry, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, entry := range s.items.EntryMap {
		if entry.Name == name {
			return entry, true
		}
	}

	return types.TodoEntry{}, false
}

func (s *ItemStore) PrintEntries() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Println("Current To-Do List:")
	fmt.Println()
	for _, entry := range s.items.EntryMap {
		fmt.Println(entry.Name)
	}
	fmt.Println()
}

func (s *ItemStore) Save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return WriteDB(s)
}
