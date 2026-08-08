package types

type TodoEntry struct {
	Name        string
	DoneStatues bool
}

type TodoList struct {
	Entries []TodoEntry
}
