package store

import (
	"os"
	"path/filepath"
	"testing"
)

const testDir = "ToDoList"
const testFile = testDir + "/todolist.json"

func chdirTemp(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, testDir), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
}

func writeTestFile(t *testing.T, content string) {
	t.Helper()
	if err := os.WriteFile(testFile, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestReadDB_Valid(t *testing.T) {
	chdirTemp(t)
	writeTestFile(t, `{"Entries":[{"Name":"alpha","DoneStatus":false}]}`)

	s := ReadDB()
	if s == nil {
		t.Fatal("ReadDB returned nil for a valid file")
	}
	entries := s.GetEntries().Entries
	if len(entries) != 1 || entries[0].Name != "alpha" {
		t.Errorf("got %v, want one entry named alpha", entries)
	}
}

func TestReadDB_FileNotFound(t *testing.T) {
	chdirTemp(t)

	if s := ReadDB(); s != nil {
		t.Errorf("ReadDB = %v, want nil for a missing file", s)
	}
}

func TestReadDB_InvalidJSON(t *testing.T) {
	chdirTemp(t)
	writeTestFile(t, `this is not json`)

	if s := ReadDB(); s != nil {
		t.Errorf("ReadDB = %v, want nil for invalid json", s)
	}
}

func TestWriteDB(t *testing.T) {
	chdirTemp(t)
	s := NewItemStore()
	s.AddEntry(CreateEntry("gamma"))
	s.Save()

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"Entries":[{"Name":"gamma","DoneStatus":false}]}`
	if string(data) != want {
		t.Errorf("file content = %s, want %s", data, want)
	}
}

func TestRoundTrip(t *testing.T) {
	chdirTemp(t)
	s := NewItemStore()
	s.AddEntry(CreateEntry("alpha"))
	s.AddEntry(CreateEntry("beta"))
	s.Save()

	loaded := ReadDB()
	if loaded == nil {
		t.Fatal("ReadDB returned nil after a successful Save")
	}
	entries := loaded.GetEntries().Entries
	if len(entries) != 2 {
		t.Fatalf("loaded %d entries, want 2", len(entries))
	}
	if entries[0].Name != "alpha" || entries[1].Name != "beta" {
		t.Errorf("round trip got %v, want [alpha beta]", entries)
	}
}
