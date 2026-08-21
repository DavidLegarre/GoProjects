package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"todolist/store"
	"todolist/types"

	"github.com/google/uuid"
)

func entryNames(entries map[uuid.UUID]types.TodoEntry) map[string]bool {
	names := map[string]bool{}
	for _, e := range entries {
		names[e.Name] = true
	}
	return names
}

func newTestArgs(t *testing.T, input string, s *store.ItemStore) *CommandArguments {
	t.Helper()
	return &CommandArguments{
		scanner:   bufio.NewScanner(strings.NewReader(input)),
		itemStore: s,
	}
}

func chdirTemp(t *testing.T) {
	t.Helper()
	t.Chdir(t.TempDir())
}

func TestGetDefaultDataPath(t *testing.T) {
	chdirTemp(t)
	got, err := getDefaultDataPath()
	if err != nil {
		t.Fatalf("getDefaultDataPath returned error: %v", err)
	}
	cwd, _ := os.Getwd()
	want := filepath.Join(cwd, "TodoList.json")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestLoadStoreFresh(t *testing.T) {
	chdirTemp(t)
	s, err := loadStore()
	if err != nil {
		t.Fatalf("loadStore returned error: %v", err)
	}
	if s == nil {
		t.Fatal("loadStore returned nil store")
	}
	if got := len(s.GetEntries().EntryMap); got != 0 {
		t.Errorf("got %d entries, want 0", got)
	}
}

func TestLoadStoreExisting(t *testing.T) {
	chdirTemp(t)
	content := `{"EntryMap":{"e0e0e0e0-e0e0-4e0e-8e0e-e0e0e0e0e0e0":{"Name":"alpha","DoneStatus":false}}}`
	if err := os.WriteFile("TodoList.json", []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	s, err := loadStore()
	if err != nil {
		t.Fatalf("loadStore returned error: %v", err)
	}
	names := entryNames(s.GetEntries().EntryMap)
	if len(names) != 1 || !names["alpha"] {
		t.Errorf("got %v, want one entry named alpha", names)
	}
}

func TestFetch(t *testing.T) {
	got, err := fetch("", bufio.NewScanner(strings.NewReader("hello\n")))
	if err != nil {
		t.Fatalf("fetch returned error: %v", err)
	}
	if got != "hello" {
		t.Errorf("got %q, want %q", got, "hello")
	}
}

func TestFetchEOF(t *testing.T) {
	_, err := fetch("", bufio.NewScanner(strings.NewReader("")))
	if !errors.Is(err, io.EOF) {
		t.Errorf("expected io.EOF, got %v", err)
	}
}

func TestCmdAdd(t *testing.T) {
	s := store.NewItemStore(filepath.Join(t.TempDir(), "todolist.json"))
	if err := cmdAdd(newTestArgs(t, "buy milk\n", s)); err != nil {
		t.Fatalf("cmdAdd returned error: %v", err)
	}
	names := entryNames(s.GetEntries().EntryMap)
	if len(names) != 1 || !names["buy milk"] {
		t.Errorf("got %v, want one entry named buy milk", names)
	}
}

func TestCmdAddEmptyName(t *testing.T) {
	s := store.NewItemStore(filepath.Join(t.TempDir(), "todolist.json"))
	if err := cmdAdd(newTestArgs(t, "   \n", s)); err != nil {
		t.Fatalf("cmdAdd returned error: %v", err)
	}
	if got := len(s.GetEntries().EntryMap); got != 0 {
		t.Errorf("got %d entries, want 0", got)
	}
}

func TestCmdAddDuplicate(t *testing.T) {
	s := store.NewItemStore(filepath.Join(t.TempDir(), "todolist.json"))
	if err := cmdAdd(newTestArgs(t, "buy milk\n", s)); err != nil {
		t.Fatalf("first cmdAdd: %v", err)
	}
	if err := cmdAdd(newTestArgs(t, "buy milk\n", s)); err != nil {
		t.Fatalf("second cmdAdd: %v", err)
	}
	if got := len(s.GetEntries().EntryMap); got != 1 {
		t.Errorf("got %d entries, want 1 (duplicate rejected)", got)
	}
}

func TestCmdAddFetchError(t *testing.T) {
	s := store.NewItemStore(filepath.Join(t.TempDir(), "todolist.json"))
	if err := cmdAdd(newTestArgs(t, "", s)); !errors.Is(err, io.EOF) {
		t.Errorf("expected io.EOF, got %v", err)
	}
	if got := len(s.GetEntries().EntryMap); got != 0 {
		t.Errorf("got %d entries, want 0", got)
	}
}

func TestCmdRemove(t *testing.T) {
	s := store.NewItemStore(filepath.Join(t.TempDir(), "todolist.json"))
	s.AddEntry(store.CreateEntry("alpha"))
	s.AddEntry(store.CreateEntry("beta"))
	if err := cmdRemove(newTestArgs(t, "alpha\n", s)); err != nil {
		t.Fatalf("cmdRemove returned error: %v", err)
	}
	names := entryNames(s.GetEntries().EntryMap)
	if len(names) != 1 || !names["beta"] {
		t.Errorf("got %v, want [beta]", names)
	}
}

func TestCmdRemoveNotFound(t *testing.T) {
	s := store.NewItemStore(filepath.Join(t.TempDir(), "todolist.json"))
	s.AddEntry(store.CreateEntry("alpha"))
	err := cmdRemove(newTestArgs(t, "zzz\n", s))
	if !errors.Is(err, store.ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
	if got := len(s.GetEntries().EntryMap); got != 1 {
		t.Errorf("got %d entries, want 1", got)
	}
}

func TestCmdList(t *testing.T) {
	s := store.NewItemStore("")
	s.AddEntry(store.CreateEntry("alpha"))
	if err := cmdList(newTestArgs(t, "", s)); err != nil {
		t.Fatalf("cmdList returned error: %v", err)
	}
}

func TestLoop(t *testing.T) {
	s := store.NewItemStore(filepath.Join(t.TempDir(), "todolist.json"))
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = old }()

	if _, err := w.WriteString("a\nbuy milk\ne\n"); err != nil {
		t.Fatal(err)
	}
	w.Close()

	if err := loop(s); err != nil {
		t.Fatalf("loop returned error: %v", err)
	}
	names := entryNames(s.GetEntries().EntryMap)
	if len(names) != 1 || !names["buy milk"] {
		t.Errorf("got %v, want one entry named buy milk", names)
	}
}