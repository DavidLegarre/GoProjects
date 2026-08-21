package store

import (
	"errors"
	"testing"
)

func TestNewItemStore(t *testing.T) {
	s := NewItemStore("")
	if s == nil {
		t.Fatal("NewItemStore returned nil")
	}
	if got := s.GetEntries().EntryMap; got == nil {
		t.Errorf("expected initialized map, got nil")
	}
}

func TestAddEntry(t *testing.T) {
	s := NewItemStore("")
	names := []string{"first", "second", "third"}
	for _, name := range names {
		s.AddEntry(CreateEntry(name))
	}

	entries := s.GetEntries().EntryMap
	if len(entries) != len(names) {
		t.Fatalf("expected %d entries, got %d", len(names), len(entries))
	}
	got := map[string]bool{}
	for _, e := range entries {
		got[e.Name] = true
	}
	for _, name := range names {
		if !got[name] {
			t.Errorf("missing entry %q", name)
		}
	}
}

func TestRemoveEntryByName(t *testing.T) {
	tests := []struct {
		name     string
		setup    []string
		remove   string
		wantErr  error
		wantLeft []string
	}{
		{name: "removes middle element", setup: []string{"a", "b", "c"}, remove: "b", wantErr: nil, wantLeft: []string{"a", "c"}},
		{name: "removes first element", setup: []string{"a", "b"}, remove: "a", wantErr: nil, wantLeft: []string{"b"}},
		{name: "removes only element", setup: []string{"a"}, remove: "a", wantErr: nil, wantLeft: nil},
		{name: "removes all duplicates", setup: []string{"a", "b", "b"}, remove: "b", wantErr: nil, wantLeft: []string{"a"}},
		{name: "missing name is untouched", setup: []string{"a", "b"}, remove: "z", wantErr: ErrNotFound, wantLeft: []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewItemStore("")
			for _, n := range tt.setup {
				s.AddEntry(CreateEntry(n))
			}

			got := s.RemoveEntryByName(tt.remove)
			if !errors.Is(got, tt.wantErr) {
				t.Errorf("RemoveEntryByName(%q) = %v, want %v", tt.remove, got, tt.wantErr)
			}

left := s.GetEntries().EntryMap
		if len(left) != len(tt.wantLeft) {
			t.Fatalf("left with %d entries, want %d", len(left), len(tt.wantLeft))
		}
		names := map[string]bool{}
		for _, e := range left {
			names[e.Name] = true
		}
		for _, n := range tt.wantLeft {
			if !names[n] {
				t.Errorf("missing %q in left entries", n)
			}
		}
	})
	}
}

func TestCreateEntry(t *testing.T) {
	e := CreateEntry("task")
	if e.Name != "task" {
		t.Errorf("Name = %q, want %q", e.Name, "task")
	}
	if e.DoneStatus {
		t.Error("expected DoneStatus to be false")
	}
}

func TestSearchEntryByName(t *testing.T) {
	s := NewItemStore("")
	s.AddEntry(CreateEntry("alpha"))
	s.AddEntry(CreateEntry("beta"))

	e, found := s.SearchEntryByName("alpha")
	if !found || e.Name != "alpha" {
		t.Errorf("SearchEntryByName(alpha) = (%v, %v), want (alpha, true)", e.Name, found)
	}

	e, found = s.SearchEntryByName("beta")
	if !found || e.Name != "beta" {
		t.Errorf("SearchEntryByName(beta) = (%v, %v), want (beta, true)", e.Name, found)
	}

	e, found = s.SearchEntryByName("missing")
	if found || e.Name != "" {
		t.Errorf("SearchEntryByName(missing) = (%v, %v), want (\"\", false)", e.Name, found)
	}
}
