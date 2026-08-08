package types

type TodoEntry struct {
	Name       string
	DoneStatus bool
}

type TodoList struct {
	Entries []TodoEntry
}
