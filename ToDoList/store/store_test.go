package store

import (
	"testing"
)

func TestNewItemStore(t *testing.T) {
	s := NewItemStore()
	if s == nil {
		t.Fatal("NewItemStore returned nil")
	}
	if got := s.GetEntries().Entries; got != nil {
		t.Errorf("expected nil entries, got %v", got)
	}
}

func TestAddEntry(t *testing.T) {
	s := NewItemStore()
	names := []string{"first", "second", "third"}
	for _, name := range names {
		s.AddEntry(CreateEntry(name))
	}

	entries := s.GetEntries().Entries
	if len(entries) != len(names) {
		t.Fatalf("expected %d entries, got %d", len(names), len(entries))
	}
	for i, name := range names {
		if entries[i].Name != name {
			t.Errorf("entry %d: expected %q, got %q", i, name, entries[i].Name)
		}
	}
}

func TestRemoveEntryByName(t *testing.T) {
	tests := []struct {
		name      string
		setup     []string
		remove    string
		wantFound bool
		wantLeft  []string
	}{
		{name: "removes middle element", setup: []string{"a", "b", "c"}, remove: "b", wantFound: true, wantLeft: []string{"a", "c"}},
		{name: "removes first element", setup: []string{"a", "b"}, remove: "a", wantFound: true, wantLeft: []string{"b"}},
		{name: "removes only element", setup: []string{"a"}, remove: "a", wantFound: true, wantLeft: nil},
		{name: "missing name is untouched", setup: []string{"a", "b"}, remove: "z", wantFound: false, wantLeft: []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewItemStore()
			for _, n := range tt.setup {
				s.AddEntry(CreateEntry(n))
			}

			got := s.RemoveEntryByName(tt.remove)
			if got != tt.wantFound {
				t.Errorf("RemoveEntryByName(%q) = %v, want %v", tt.remove, got, tt.wantFound)
			}

			left := s.GetEntries().Entries
			if len(left) != len(tt.wantLeft) {
				t.Fatalf("left with %d entries, want %d", len(left), len(tt.wantLeft))
			}
			for i, n := range tt.wantLeft {
				if left[i].Name != n {
					t.Errorf("left[%d] = %q, want %q", i, left[i].Name, n)
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
	s := NewItemStore()
	s.AddEntry(CreateEntry("alpha"))
	s.AddEntry(CreateEntry("beta"))

	found := s.SearchEntryByName("alpha")
	if found == nil {
		t.Fatal("expected to find alpha")
	}
	if found.Name != "alpha" {
		t.Errorf("Name = %q, want %q", found.Name, "alpha")
	}

	if got := s.SearchEntryByName("missing"); got != nil {
		t.Errorf("expected nil for missing name, got %v", got)
	}

	if got := s.SearchEntryByName("alpha"); got != &s.items.Entries[0] {
		t.Error("returned pointer is not the stored entry")
	}
}
