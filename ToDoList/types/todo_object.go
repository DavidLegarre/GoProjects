package types

import (
	"github.com/google/uuid"
)

type TodoEntry struct {
	Id         uuid.UUID
	Name       string
	DoneStatus bool
}

type TodoList struct {
	EntryMap map[uuid.UUID]TodoEntry
}
