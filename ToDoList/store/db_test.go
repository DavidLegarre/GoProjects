package store

import (
	"errors"
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

	s, err := ReadDB(testFile)
	if err != nil {
		t.Fatalf("ReadDB returned unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("ReadDB returned a nil store")
	}
	entries := s.GetEntries().Entries
	if len(entries) != 1 || entries[0].Name != "alpha" {
		t.Errorf("got %v, want one entry named alpha", entries)
	}
}

func TestReadDB_FileNotFound(t *testing.T) {
	chdirTemp(t)

	_, err := ReadDB(testFile)
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected os.ErrNotExist, got %v", err)
	}
}

func TestReadDB_InvalidJSON(t *testing.T) {
	chdirTemp(t)
	writeTestFile(t, `this is not json`)

	_, err := ReadDB(testFile)
	if err == nil {
		t.Fatal("expected an error for invalid json")
	}
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("invalid json must not be a not-exist error, got %v", err)
	}
}

func TestReadDB_PathIsDirectory(t *testing.T) {
	chdirTemp(t)
	if err := os.Mkdir(testFile, 0o755); err != nil {
		t.Fatal(err)
	}

	_, err := ReadDB(testFile)
	if err == nil {
		t.Fatal("expected an error when the path is a directory")
	}
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("a directory must not be a not-exist error, got %v", err)
	}
}

func TestWriteDB(t *testing.T) {
	chdirTemp(t)
	s := NewItemStore(testFile)
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
	matches, err := filepath.Glob(filepath.Join(testDir, "tmp.*.json"))
	if err != nil || len(matches) != 0 {
		t.Errorf("stale tmp file left behind: %v", matches)
	}
}

func TestRoundTrip(t *testing.T) {
	chdirTemp(t)
	s := NewItemStore(testFile)
	s.AddEntry(CreateEntry("alpha"))
	s.AddEntry(CreateEntry("beta"))
	s.Save()

	loaded, err := ReadDB(testFile)
	if err != nil {
		t.Fatalf("ReadDB returned unexpected error: %v", err)
	}
	if loaded == nil {
		t.Fatal("ReadDB returned a nil store")
	}
	entries := loaded.GetEntries().Entries
	if len(entries) != 2 {
		t.Fatalf("loaded %d entries, want 2", len(entries))
	}
	if entries[0].Name != "alpha" || entries[1].Name != "beta" {
		t.Errorf("round trip got %v, want [alpha beta]", entries)
	}
}

func TestRoundTripResaveAfterLoad(t *testing.T) {
	chdirTemp(t)
	s := NewItemStore(testFile)
	s.AddEntry(CreateEntry("alpha"))
	s.Save()

	loaded, err := ReadDB(testFile)
	if err != nil {
		t.Fatalf("ReadDB returned unexpected error: %v", err)
	}
	loaded.AddEntry(CreateEntry("beta"))
	if err := loaded.Save(); err != nil {
		t.Fatalf("Save after load failed: %v", err)
	}

	reloaded, err := ReadDB(testFile)
	if err != nil {
		t.Fatalf("ReadDB returned unexpected error: %v", err)
	}
	if got := len(reloaded.GetEntries().Entries); got != 2 {
		t.Errorf("reloaded %d entries, want 2", got)
	}
}
